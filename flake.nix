{
  description = "Email logging service for ServeUp";

  inputs = rec {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-22.05";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }: utils.lib.eachDefaultSystem (system:
    let
      version = "0.1.0";
      pkgs = import nixpkgs { inherit system; };
      #olderpkgs = import oldnixpkgs { inherit system; };

      # TODO
      # An older package repo is used to pull version 6.2.0 of QEMU
      # This version of QEMU is required for Podmand to work correctly
      # on MacOS Monterey
      #olderPkgs = import (builtins.fetchGit {
      #  name = "revision_for_qemu6.2.0";                                                 
      #  url = "https://github.com/NixOS/nixpkgs/";                       
      #  ref = "refs/heads/nixos-unstable";                     
      #  rev = "d1c3fea7ecbed758168787fe4e4a3157e52bc808";                                           
      #}) {};                                                                           

      #qemu-6_2_0 = olderPkgs.qemu_full;
    in with pkgs; rec {
      packages = utils.lib.flattenTree {
        main = buildGoModule {
          inherit version;
          pname = "emailog";

          vendorSha256 = "sha256-OVu4XvjUrqNR6yVsp56xEAXXSZ/UAVYCCVwSG+lGTSw=";#"sha256-OVu4XvjUrqNR6yVsp56xEAXXSZ/UAVYCCVwSG+lGTSw=";

          src = ./.;

          doCheck = false;

          meta = {
            description = "Email logging HTTP server, written in Go";
            homepage = "https://github.com/ServeUp-Inc/emailog";
            license = lib.licenses.mit;
            maintainers = [ "justkash" ];
            platforms = lib.platforms.linux ++ lib.platforms.darwin;
          };
        };
      };

      defaultPackage = packages.main;
      defaultApp = packages.main;

      devShells.default = mkShell rec {
        packages = [
          #qemu-6_2_0
        ];
        buildInputs = [
          gnupg
          go
          python310Packages.grip
          podman
          curl
          git
        ];
        SERVER_HOST = "127.0.0.1";
        SERVER_PORT = 4000;
        DB_CONTAINER_NAME = "devsqldb";
        DB_CONTAINER_PORT = 3306;
        DB_ROOT_PASS = "devroot";
        DB_USER = "devuser";
        DB_PASS = "devpass";
        DB_NAME = "devdb";
        DB_PORT = DB_CONTAINER_PORT;
        DB_HOST = "127.0.0.1";
        shellHook = ''
          start_markdown_server() {
            grip > /dev/null 2>&1 &
          }

          stop_markdown_server() {
            # List jobs, find grip, extract process id, pass to kill command
            jobs -l | grep grip | awk '{print $2}' | xargs kill
          }

          start_podman() {
            echo 'Starting Podman...'
            podman machine init > /dev/null 2>&1
            podman machine start > /dev/null 2>&1

            # Wait for the default machine to be ready
            #until [ "`podman machine inspect --format {{.State}}`"=="running" ]; do
            #  sleep 0.1;
            #done;

            # Wait 5s for machine to actually start
            # Podman machine state seems to change to running
            # before actually being ready
            # TODO Change to better check
            sleep 15s

            echo 'Podman started'
          }

          stop_podman() {
            podman machine stop 
            echo 'Podman stopped'
          }

          start_db_container() {
            echo 'Starting database container...'

            # Run MySQL instance, tag is oracle because it supports arm64/v8
            podman run -d \
              --name $DB_CONTAINER_NAME \
              -p 3306:$DB_CONTAINER_PORT \
              -e MYSQL_ROOT_PASSWORD=$DB_ROOT_PASS \
              -e MYSQL_USER=$DB_USER \
              -e MYSQL_PASSWORD=$DB_PASS \
              -e MYSQL_DATABASE=$DB_NAME \
              mysql:oracle > /dev/null

            # Wait for container to be ready
            until [ "`podman inspect -f {{.State.Health.Status}} $DB_CONTAINER_NAME`"=="healthy" ]; do
              sleep 0.1;
            done;

            echo 'Database container ready'
          }

          stop_db_container() {
            podman container stop $DB_CONTAINER_NAME > /dev/null
            podman container wait $DB_CONTAINER_NAME > /dev/null
            podman container prune --force > /dev/null
            echo 'Database container stopped.'
          }

          setup() {
            start_podman
            start_db_container
            start_markdown_server
          }

          teardown() {
            stop_db_container
            stop_podman
            stop_markdown_server
          }

          setup

          trap 'teardown' EXIT
        '';
      };
    }
  );
}
