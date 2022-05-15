{
  description = "Parser for bio scientific names";
  inputs.nixpkgs.url = github:NixOS/nixpkgs;

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
    lib = pkgs.lib;
  in {
    defaultPackage.${system} = pkgs.callPackage ./default.nix {};
    devShell.${system} = pkgs.callPackage ./shell.nix {};
  };
}