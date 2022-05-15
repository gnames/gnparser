{ mkShell, go, gopls }:
mkShell rec {
  buildInputs = [ go gopls ];
}
