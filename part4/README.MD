CSI2120[A] <br>
FEIYU LIN #300298455 <br>
CHRISTOPHER LIT #300298516 <br>

## Getting Started: PROJECT PART 4
NOTE: the dataset are NOT provided in the zip file but the pre-computed query images with histograms are provided

### to compile and run:
inside [similaritySearch.pl](similaritySearch.pl), find the code
```prolog
dataset('C:\\Users\\linfe\\OneDrive\\Desktop\\csi2120\\project\\part1\\imageDataset2_15_20').
```
you MUST change this to the directory of the dataset <br><br>

open SWI-PROLOG and consult [similaritySearch.pl](similaritySearch.pl)<br>
you should see compilation success similar as the following in the terminal:
```bash
% c:/Users/linfe/OneDrive/Desktop/csi2120/project/part4/similaritySearch.pl compiled 0.00 sec, 13 clauses
```
now run the following predicate: 
```bash
similarity_search('<full filepath to your text file containing pre-computed histograms of the query image>', SimilarList).
```

FOR EXAMPLE:
since I put my query images folder and dataset folder in part1 of this project, I would need the full filepath to that directory
```bash
similarity_search('C:/Users/linfe/OneDrive/Desktop/csi2120/project/part4/queryImagesWithHistograms/queryImages/q00.jpg.txt', SimilarList).
```





