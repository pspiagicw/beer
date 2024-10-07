{
  description = "Go development environment";

  inputs = {
    # Nixpkgs repository
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
    custom.url = "github:pspiagicw/packages";
  };

  outputs = {
    self,
    nixpkgs,
    custom,
  }: {
    devShell.x86_64-linux = let
      pkgs = import nixpkgs {system = "x86_64-linux";};
    in
      pkgs.mkShell {
        buildInputs = [
          pkgs.gopls
          pkgs.delve
          pkgs.gotools
          pkgs.go
        ];
      };
    devShell.aarch64-linux = let
      pkgs = import nixpkgs {system = "aarch64-linux";};
    in
      pkgs.mkShell {
        buildInputs = [
          pkgs.gopls
          pkgs.delve
          pkgs.gotools
          pkgs.go
          custom.packages.aarch64-linux.groom
        ];
      };
  };
}
