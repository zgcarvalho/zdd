acceptors = ['N.1',  'N.ar', 'O.2', 'O.co2']
donors = [ 'N.4', 'N.am', 'N.pl3', 'P.3']
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
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.1,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.01,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in donors:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.1,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.01,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.1,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in aAndD:
            if atom[j] in acceptors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #H-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.8,4.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,2.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in halogens:
            if atom[j] in acceptors:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.01,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,2.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Halogen-bond
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,2.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Halogen-bond (repulsive?)
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.01,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in metals:
            if atom[j] in acceptors:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.1,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #Metal
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],2.4,7.0,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #repulsive
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.2,-0.1,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        elif atom[i] in nonpolar:
            if atom[j] in acceptors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in donors:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in aAndD:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in halogens:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in metals:
                #buried
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.05,1.0,20.0,1.0,1.0,1.0)
            elif atom[j] in nonpolar:
                #non-polar
                print "{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}\t{}".format(atom[i], atom[j],3.6,0.4,1.0,20.0,1.0,1.0,1.0)
            else:
                print "Problem: atom not founded"
        else:
            print "Problem: atom not founded"
