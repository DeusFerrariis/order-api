{
  description = "Development shell for Go development.";

  inputs = {
    nixpkgs.url      = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url  = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {};
      in
      with pkgs;
      {
        devShells.default = mkShell {
          buildInputs = [
            go
            gopls
          ];
        };
      }
    );
}
