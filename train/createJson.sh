IFS=$'\n'
for i in `cat pkd`; do
  name=$(echo $i | cut -f1 -d$'\t')
  pkd=$(echo $i | cut -f2 -d$'\t')
  printf "{\"name\":\"${name}\",\"pkd\":${pkd},\"receptor\":\"${name}_protein.mol2\",\"true\":\"ligand/${name}_ligand.mol2\", \"falses\":["
  for j in `ls falses/$name*`; do
    printf \"$j\",
  done
  printf "]}\n"
done
