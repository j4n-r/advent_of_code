let
  pkgs = import <nixpkgs> {
    config = { };
    overlays = [ ];
  };
  lib = pkgs.lib;
  content = lib.strings.fileContents ./test_input.txt;
  cleaned = lib.strings.splitString "," (lib.strings.replaceString "\n" "" content);
  id_pairs = map (pair: map lib.strings.toInt (lib.strings.splitString "-" pair)) cleaned;

  checkPair = (
    pair:
    let
      first = lib.debug.traceVal (builtins.elemAt pair 0);
      second = lib.debug.traceVal (builtins.elemAt pair 1);
      # first = builtins.elemAt pair 0;
      # second = builtins.elemAt pair 1;
    in

      lib.debug.traceVal (first == second) 
  );
in
map (pair:  (checkPair pair)) id_pairs
