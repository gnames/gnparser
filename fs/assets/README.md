# Vocabularies

`bacteria_genera.txt`
: this list is used to mark parsed names as bacterial names.

`bacteria_genera_homonyms.txt`
: this list contains bacterial generic names that exist under ICZN or ICN codes.

`genera_auth_icn.txt`
: this list contains authors of genera under ICN codes.

## Creation of genera_auth_icn.txt

1. Get the latest IRMNG file.
2. Extract authors of ICN genera
3. Parse the authors and take only "basionym" authors (makes list 500 authors smaller)
4. Break authors to words, collect words that are capitalized, have no periods, larger than 2 characters.
5. Clean up authors from spaces, commas, parentheses.
6. Create list of all genera (canonical form)
7. Remove from authors list all genera names.