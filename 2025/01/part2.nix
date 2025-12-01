let
  pkgs = import <nixpkgs> {
    config = { };
    overlays = [ ];
  };
  lib = pkgs.lib;
  contents = lib.fileContents ./test_input.txt;
  # contents = lib.fileContents ./input.txt;

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

  move_dial_left = (
    {
      curr,
      op_val,
      count,
    }:
    if lib.debug.traceSeq {inherit curr op_val count;} curr - op_val < 0 then
    # if  curr - op_val <= 0 then
      move_dial_left {
        count = count + 1;
        curr = (curr - op_val) + 100;
        op_val = 0;
      }
    else
      {
        count = if curr == 0 then count+1 else count;
        dial_pos = curr - op_val;
      }
  );

  move_dial_right = (
    {
      curr,
      op_val,
      count,
    }:
    if lib.debug.traceSeq {inherit curr op_val count;} curr + op_val > 100 then
    #if curr + op_val > 100 then
      move_dial_right {
        count = count + 1;
        curr = (curr + op_val) - 100;
        op_val = 0;
      }
    else
      {
        count = if curr == 0 then count+1 else count;
        dial_pos = curr + op_val;
      }
  );
  # returns the new dial pos
  getNewValue = (
    { curr_dial_pos, op, count }:
    let
      # curr_dial_pos_dbg = lib.debug.traceSeq {
      #   inherit op;
      #   curr_val = curr_dial_pos;
      # } curr_dial_pos;
    in
    if op.op == "L" then
      # if builtins.trace (curr_val - op.val < 0) curr_val - op.val < 0 then
      move_dial_left {count=count; curr=curr_dial_pos; op_val = (op.val);}
    else
      move_dial_right {count=count; curr=curr_dial_pos; op_val = (op.val);}
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
        res = getNewValue {
          inherit op count;
          curr_dial_pos = dial_pos;
        };
        new_pos = res.dial_pos;
        new_count = res.count;
      in
        check_list {
          count = new_count;
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
