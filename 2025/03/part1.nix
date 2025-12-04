let
  pkgs = import <nixpkgs> {
    config = { };
    overlays = [ ];
  };
  lib = pkgs.lib;
  input = lib.strings.splitString "\n" (lib.strings.fileContents ./test_input.txt);
  battery_banks_str = map (bank: (lib.strings.stringToCharacters bank)) input;
  battery_banks = map (split_bank: map (char: lib.strings.toInt char) split_bank) battery_banks_str;
in
battery_banks


