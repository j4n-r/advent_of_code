let
  pkgs = import <nixpkgs> {
    config = { };
    overlays = [ ];
  };
  lib = pkgs.lib;
  content = lib.strings.fileContents ./input.txt;
  cleaned = lib.strings.splitString "," (lib.strings.replaceString "\n" "" content);
  id_pairs = map (pair: (lib.strings.splitString "-" pair)) cleaned;

  is_valid = (
    string:
    let
      len = builtins.stringLength string;
      half_len = builtins.div len 2;
      first_half = builtins.substring 0 half_len string;
      second_half = builtins.substring half_len len string;
    in
    if lib.mod len 2 != 0 then true else first_half != second_half
  );
  check_range = (
    list:
    let
      start = lib.strings.toInt (builtins.elemAt list 0);
      finish = lib.strings.toInt (builtins.elemAt list 1);
      ids = lib.lists.range start finish;
    in
    builtins.foldl' (
      acc: id:
      let
        id_str = builtins.toString id;
      in
      if !(is_valid id_str) then (acc + id) else acc + 0
    ) 0 ids
  );
in
builtins.foldl' (
  (acc:
  range:
  acc + (check_range range))
) 0 id_pairs
