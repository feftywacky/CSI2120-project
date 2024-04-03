/*
 * CSI2120[A]
 * FEIYU LIN #300298455
 * CHRISTOPHER LIT #300298516
 */


import java.io.File;
import java.util.PriorityQueue;
import java.util.Map;
import java.util.HashMap;
import java.util.Comparator;


public class SimilaritySearch {
    public static void main(String[] args) {

        System.out.println("The query image is: " + args[0]);
        System.out.println("The dataset is: " + args[1]);
        System.out.println("Searching for similar images...");

        String query_image_filePath = "queryImages/";
        query_image_filePath += args[0];
        
        String dataset_filePath = args[1];
        dataset_filePath += "/";
        File image_dataset_filePath = new File(dataset_filePath);

        // assume that the histograms of the image dataset have been pre-computed

        ColorImage query_image = new ColorImage(query_image_filePath);

        // we want to hisotorgram to have 3-bit color space
        ColorHistogram query_image_hist = new ColorHistogram(3);

        // note that setImage will reduce the color space of the image to d-bit
        query_image_hist.setImage(query_image); 

        // save the histogram of the query image
        query_image_hist.save(args[0] + ".txt");

        // init priority queue with custom comparator to filter 5 most similar images
        // maps the filename to the similarity score and compares it by the similarity score (retrived using lambda function)
        PriorityQueue<Map.Entry<String, Double>> pq = new PriorityQueue<>(Comparator.comparing((entry) -> entry.getValue()));

        // compare the query image with each image in the dataset and keep track of the 5 most similar images using a priority queue
        for (File file : image_dataset_filePath.listFiles()) {
            if (file.isFile() && file.getName().endsWith(".txt")) {
                ColorHistogram dataset_hist = new ColorHistogram(dataset_filePath + file.getName());
                double similarity = query_image_hist.compare(dataset_hist);
                
                pq.offer(new HashMap.SimpleEntry<>(file.getName(), similarity));
                
                // ensure only 5 most similar images are in the priority queue
                if (pq.size() > 5) {
                    pq.poll();
                }

            }
        }

        // print out the 5 most similar images from the dataset to the query image
        int count = 5;
        System.out.println("");
        System.out.println("The 5 most similar images from the dataset to " + args[0] + " are: ");
        while (!pq.isEmpty()) {
            Map.Entry<String, Double> entry = pq.poll();
            System.out.println("#" + count + " => Image: " + entry.getKey() + ", Similarity: " + entry.getValue());
            count--;
        }
    }
}