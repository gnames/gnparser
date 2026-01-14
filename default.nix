{
  lib,
  buildGoModule,
  fetchFromGitHub,
  stdenv,
  glibc,
}:
buildGoModule rec {
  pname = "gnparser";
  version = "v1.14.1";
  date = "2026-01-14";

  src = lib.cleanSourceWith {
    filter = name: type: let
      baseName = baseNameOf (toString name);
    in
      !(lib.hasInfix "/vendor" name || baseName == "vendor");
    src = lib.cleanSource ./.;
  };

  vendorHash = "sha256-Oitse5Te35Cs8Ub7ueDw60JM5p8FOsyMsZif8OQQ+uE=";

  buildInputs = [
    stdenv
    glibc.static
  ];

  doChecks = false;

  subPackages = "gnparser";

  buildFlags = ["-mod=readonly"];

  ldflags = [
    "-s"
    "-w"
    "-linkmode external"
    "-extldflags"
    "-static"
    "-X github.com/gnames/gnparser.Version=${version}"
    "-X github.com/gnames/gnparser.Build=${date}"
  ];

  meta = with lib; {
    description = "Parser for bio scientific names";
    homepage = "https://github.com/gnames/gnparser";
    license = licenses.mit;
    maintainers = with maintainers; ["dimus"];
  };
}
