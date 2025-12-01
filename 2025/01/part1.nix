let
  pkgs = import <nixpkgs> {
    config = { };
    overlays = [ ];
  };
  lib = pkgs.lib;
  # contents = lib.fileContents ./test_input.txt;
  contents = lib.fileContents ./input.txt;

  splits = lib.splitString "\n" contents;

  ops = map (
    op:
    if lib.strings.hasPrefix "L" op then
      {
        op = "L";
        val = lib.strings.toInt (lib.strings.removePrefix "L" op);
      }
    else
      {
        op = "R";
        val = lib.strings.toInt (lib.strings.removePrefix "R" op);
      }
  ) splits;

  move_dial = curr: op: lib.mod (curr + op) 100;

  # returns the new dial pos
  getNewValue = (
    { curr_dial_pos, op }:
    let
      curr_dial_pos_dbg = lib.debug.traceSeq {
        inherit op;
        curr_val = curr_dial_pos;
      } curr_dial_pos;
    in
    if op.op == "L" then
      # if builtins.trace (curr_val - op.val < 0) curr_val - op.val < 0 then
      move_dial curr_dial_pos_dbg (-op.val)
    else
      move_dial curr_dial_pos_dbg op.val
  );

  check_list = (
    {
      count,
      list,
      dial_pos,
    }:
    if list == [ ] then
      {
        inherit count list;
        curr_val = dial_pos;
      }
    else
      let
        # op = lib.debug.traceSeq (list) (builtins.head list);
        op = builtins.head list;
        # new_value = builtins.trace (getNewValue { inherit curr_val op; }) getNewValue {
        new_pos = getNewValue {
          inherit op;
          curr_dial_pos = dial_pos;
        };
      in
      if lib.debug.traceSeq { inherit count dial_pos; } (dial_pos == 0) then
        check_list {
          count = count + 1;
          list = builtins.tail list;
          dial_pos = new_pos;
        }
      else
        check_list {
          count = count;
          list = builtins.tail list;
          dial_pos = new_pos;
        }
  );

in
check_list {
  count = 0;
  list = ops;
  dial_pos = 50;
}
