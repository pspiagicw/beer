{
  description = "Neovim configuration with stylua";

  inputs = {
    # Nixpkgs repository
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";

    custom.url = "github:pspiagicw/flakes";
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
          pkgs.go
          pkgs.gopls
          pkgs.delve
          pkgs.gotools
          custom.packages.x86_64-linux.groom
        ];
      };
  };
}