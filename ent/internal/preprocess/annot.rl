package preprocess

import (
)

func AnnotationRL(data []byte) bool {
  %%{
    machine annot;
    write data;
  }%%

  cs, p, pe, eof := 0, 0, len(data), len(data)
  _ = eof
  _ = annot_en_main
  _ = annot_error
  _ = annot_first_final

  var match bool

  %%{
    action setMatch {match = true}
    action setPos {pos = append(pos,p)}

    notes = ("species"i | "group"i | "clade"i | "authors"i | "non" | "nec" |
      "fide" | "vide" );
    tc1 = ("sensu"i | "auct"i | "sec"i | "near" | "str") "."?;
    tc2 = "("? "s." space? ([sl] | "str" | "lat") ".";
    tc3 = "pro parte"i | "p." space? "p.";
    tc4 = "("? ("nomen"i | "nom."i | "comb.");

    main := any* ((space+ | "," space?)
            (notes | tc1 |tc2 | tc3 | tc4))  %/setMatch
            ((space | punct) >setMatch);

    write init;
    write exec;
  }%%

  return match
}
