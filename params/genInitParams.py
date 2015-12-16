acceptors = ['N.1',  'N.ar', 'O.2', 'O.co2']
donors = [ 'N.4', 'N.am', 'N.pl3']
aAndD = ['N.2', 'N.3','O.3']
halogens = ['Br','Cl', 'F', 'I']
metals = ['Al', 'Ca', 'C.cat',  'Co.oh', 'Cr.oh', 'Cr.th', 'Cu', 'Fe', 'K', 'Li', 'Mg', 'Mn', 'Mo', 'Na', 'Sn', 'Zn']
nonpolar = ['C.1', 'C.2', 'C.3','C.ar', 'P.3', 'S.2', 'S.3', 'Se', 'Si', 'S.o', 'S.o2']

with open("usedAtoms") as f:
    atom = f.read().splitlines()

print "A1\tA2\tDbest\tAlpha\tBeta\tPenal\tWa\tWb\tWpenal"
for i in range(len(atom)):
    for j in range(i,len(atom)):
        if atom[i] in acceptors:
            if atom[j] in acceptors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",3.0,-0.03,1.2,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",3.4,-0.006,0.9,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in donors:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",3.0,-0.03,1.2,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",2.5,0.3,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",3.0,-0.03,1.2,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in aAndD:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"hbond",3.3,0.4,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",2.5,0.3,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in halogens:
            if atom[j] in acceptors:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",3.4,-0.006,0.9,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",2.5,0.3,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha",2.5,0.3,1.05,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"ha-repulsive",3.4,-0.006,0.9,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in metals:
            if atom[j] in acceptors:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",3.0,-0.03,1.2,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"metal",2.5,0.72,0.6,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"repulsive",3.0,-0.03,1.2,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in nonpolar:
            if atom[j] in acceptors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            elif atom[j] in donors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            elif atom[j] in metals:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"buried",3.6,0.002,1.0,9.5,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #non-polar
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i],atom[j],"npolar",3.6,0.03,0.9,9.5,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        else:
            print "Problem: atom not founded"
