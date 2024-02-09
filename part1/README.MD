#
CSI2120[A] <br>
PROJECT PART 1 <br>
FEIYU LIN #300298455 <br>
CHRISTOPHER LIT #300298516 <br>

## Getting Started
1. ensure dataset folder is named exactly "imageDataset2_15_20"
2. ensure query image folder is named exactly "queryImages"
3. ensure the dataset and query image folders are in the same directory as the java files

the code used to parse images has a hard coded file path shown below

```java
String query_image_filePath = "queryImages/";
query_image_filePath += args[0];
        
String dataset_filePath = args[1];
dataset_filePath += "/";
File image_dataset_filePath = new File(dataset_filePath);
```

### to compile and run:
```bash
javac SimilaritySearch.java
java SimilaritySearch <query image name>.jpg imageDataset2_15_20
```

for example:
```bash
javac SimilaritySearch.java
java SimilaritySearch q00.jpg imageDataset2_15_20
```





