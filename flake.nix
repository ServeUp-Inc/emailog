{
  description = "Email logging service for ServeUp";

  inputs = rec {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-22.05";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }: utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };

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
    in {
      devShells.default = with pkgs; mkShell {
        buildInputs = [
          git
          python310Packages.grip
          #qemu-6_2_0
          podman
          go
        ];
      };
    }
  );
}
