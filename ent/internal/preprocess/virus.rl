package preprocess

func IsVirus(data []byte) bool {
  %%{
    machine virus;
    write data;
  }%%

  cs, p, pe, eof := 0, 0, len(data), len(data)
  _ = eof
  _ = virus_en_main
  _ = virus_error
  _ = virus_first_final

  var match bool

  %%{
    action setMatch {match = true}

    vir_str = (alnum* "virus"i "es"i?) |
              'ICTV' | 'Ictv' | 
              ("cyano"i | "bacterio"i | "viro"i)? "phage"i "s"i? |
              ("vector"i | "viroid"i | "particle"i | "prion"i) "s"i? |
              alnum* "npv"i |
              ("alpha"i | "beta"i)? "satellite"i "s"i?;


    main := ('' | any* (space | punct))
            vir_str %/setMatch
            ((space | punct) >setMatch);

    write init;
    write exec;

  }%%

  return match
}
