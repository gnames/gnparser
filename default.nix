{
  lib,
  buildGoModule,
  fetchFromGitHub,
  stdenv,
  glibc,
}:
buildGoModule rec {
  pname = "gnparser";
  version = "v1.12.1";
  date = "2025-12-20";

  src = lib.cleanSourceWith {
    filter = name: type: let
      baseName = baseNameOf (toString name);
    in
      !(lib.hasInfix "/vendor" name || baseName == "vendor");
    src = lib.cleanSource ./.;
  };

  vendorHash = "sha256-Yl2jBQw7UFq4djhX18/k25Bs81giwBQ1VG0y1uGI1Bc=";

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
