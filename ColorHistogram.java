import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.io.PrintWriter;

public class ColorHistogram {

    // bins is an array of doubles that will store the histogram
    private double[] bins;
    private int d;

    public ColorHistogram (int d) {
        this.bins = new double[(int) Math.pow(2, d*3)];
        this.d = d;
    }
    
    public ColorHistogram (String filename) {
        int total_pixel_count = 0;
        try
        {
            File file = new File(filename);
            Scanner scanner = new Scanner(file);
    
            if (scanner.hasNext())
            {
                int size = Integer.parseInt(scanner.next());
                this.bins = new double[size];
    
                int index = 0;
                while(scanner.hasNext())
                {
                    int pixel_count = Integer.parseInt(scanner.next());
                    total_pixel_count += pixel_count;
                    this.bins[index] = pixel_count;
                    index++;
                }
            }
            else
                System.out.println("File is empty");
    
            scanner.close();
        }
        catch (FileNotFoundException e)
        {
            e.printStackTrace();
        }

        // normalize the histogram
        for (int i = 0; i < this.bins.length; i++)
            this.bins[i] /= total_pixel_count;
    }

    public void setImage(ColorImage image)
    {
        // reset bins
        // if bins was loaded from file, it will be overwritten by this method
        for(int i=0;i<bins.length;i++)
            bins[i] = 0.0;
        
        int width = image.getWidth();
        int height = image.getHeight();

        // reduce the image color space
        image.reduceColor(this.d);

        int pixels[][][] = image.getPixels();
        int total_pixel_count = width * height;

        for (int x=0;x<pixels.length;x++)
        {
            for (int y=0;y<pixels[0].length;y++)
            {
                int rgb[] = pixels[x][y];
                int bin_index = (rgb[0]<<(2*this.d)) + (rgb[1]<<this.d) + (rgb[2]); 
                bins[bin_index]++;
            }
        }

        // normalize the histogram
        for (int i = 0; i < this.bins.length; i++)
            this.bins[i] /= total_pixel_count;
    }

    public double[] getHistogram()
    {
        return this.bins;
    }

    public double compare(ColorHistogram hist)
    {
        if (this.bins.length != hist.bins.length)
            return -1;
        
        double similarity = 0.0;

        for(int i=0;i<this.bins.length;i++)
            similarity += Math.min(this.bins[i], hist.bins[i]);

        return similarity;
    }   

    public void save(String filename)
    {
        try
        {
            PrintWriter writer = new PrintWriter(filename);
            writer.println(bins.length);
    
            for (double pixel_count : bins) {
                writer.print(pixel_count + " ");
            }
            writer.close();
        }
        catch (FileNotFoundException e)
        {
            e.printStackTrace();
        }
    }
}