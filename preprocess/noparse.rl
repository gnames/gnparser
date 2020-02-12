package preprocess

func NoParse(data []byte) bool {

  %%{
    machine noparse;
    write data;
  }%%

  cs, p, pe, eof := 0, 0, len(data), len(data)
  _ = eof
	_ = noparse_first_final
	_ = noparse_error
	_ = noparse_en_main

  var match bool


  %%{
    action setMatch {match = true}

    noparse1 = ("Not" | "None" | "Un" ("n"? "amed" | "identified"));
    noparse2 = any* [Ii] "nc" ("." | "ertae") space* [Ss] "ed" ("." | "is");
    noparse3 = any* ("phytoplasma" | "plasmid" "s"? | [^A-Z] "RNA" [^A-Z]*);


    main := (noparse1 | noparse2 | noparse3) %/setMatch
            ((space | punct) >setMatch);

    write init;
    write exec;

  }%%

  return match
}
