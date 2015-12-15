IFS=$'\n'
for i in `cat energy`; do
  name=$(echo $i | cut -f1 -d$'\t')
  energy=$(echo $i | cut -f2 -d$'\t')
  printf "{\"name\":\"${name}\",\"energy\":${energy},\"receptor\":\"receptors/${name}_protein.mol2\",\"positive\":\"ligands/${name}_ligand.mol2\", \"negatives\":["
  for j in `ls falses/$name*`; do
    printf \"$j\",
  done
  printf "]},\n"
done
