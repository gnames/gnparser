{ lib, buildGoModule, fetchFromGitHub, stdenv, glibc }:

buildGoModule rec {
  pname = "gnparser";
  version = "v1.6.6";
  date = "2022-05-17";

  src = ./.;

  vendorSha256 = "sha256-TY/vIgtu/GeVKJ1AonMMxCvIbK3ATc2jp9Zqq1YQ9Mg=";

  buildInputs = [
    stdenv
    glibc.static
  ];

  doChecks = false;

  subPackages = "gnparser";

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
    maintainers = with maintainers; [ "dimus" ];
  };
}
