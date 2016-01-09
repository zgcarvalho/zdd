acceptors = ['N.1',  'N.ar', 'O.2', 'O.co2']
donors = [ 'N.4', 'N.am', 'N.pl3']
aAndD = ['N.2', 'N.3','O.3']
halogens = ['Br','Cl', 'F', 'I']
metals = ['Al', 'Ca', 'C.cat',  'Co.oh', 'Cr.oh', 'Cr.th', 'Cu', 'Fe', 'K', 'Li', 'Mg', 'Mn', 'Mo', 'Na', 'Sn', 'Zn']
nonpolar = ['C.1', 'C.2', 'C.3','C.ar', 'P.3', 'S.2', 'S.3', 'Se', 'Si', 'S.o', 'S.o2']

opt = [1.7203856221843745, 3.841140939682709, -1.2523686486972208, 0.059696567926524724, #penal / metal
3.1100251465313025, -0.34828173005247387, 0.4029551227223057, #repulsive
 2.7812437371094925, 0.26700530529527583, 0.14649880558141376, #buried
 2.016131768252614, -0.29395532027549653, 1.3965986170459754, #hbond
 3.5249493533634104, -0.7378164951514568, 0.934863488067968, #Halogen-bond
  2.4557261653647275, 0.6640871593057627, 0.3588161094134191, #halogen-repulsive
  4.211394128713735, -0.15452766623671943, 0.6084459472500465] #non-polar

opt[0] += 3.0
with open("usedAtoms") as f:
    atom = f.read().splitlines()

print "A1\tA2\tDbest\tAlpha\tBeta\tPenal\tWa\tWb\tWpenal"
for i in range(len(atom)):
    for j in range(i,len(atom)):
        if atom[i] in acceptors:
            if atom[j] in acceptors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",opt[16],opt[17],opt[18],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in donors:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",opt[13],opt[14],opt[15],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in aAndD:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",opt[10],opt[11],opt[12],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",opt[13],opt[14],opt[15],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in halogens:
            if atom[j] in acceptors:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",opt[16],opt[17],opt[18],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",opt[13],opt[14],opt[15],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",opt[13],opt[14],opt[15],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",opt[16],opt[17],opt[18],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in metals:
            if atom[j] in acceptors:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",opt[1],opt[2],opt[3],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in nonpolar:
            if atom[j] in acceptors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in donors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in halogens:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in metals:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",opt[4],opt[5],opt[6],opt[0],1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #non-polar
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"npolar",opt[19],opt[20],opt[21],opt[0],1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        else:
            print "Problem: atom not founded"
