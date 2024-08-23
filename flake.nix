{
  description = "Neovim configuration with stylua";

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
        hardeningDisable = ["fortify"];
        buildInputs = [
          pkgs.go
          pkgs.gopls
          pkgs.delve
          pkgs.nodejs_22
          pkgs.gotools
          custom.packages.x86_64-linux.groom
        ];
      };
  };
}
