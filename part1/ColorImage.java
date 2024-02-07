/*
 * CSI2120[A]
 * FEIYU LIN #300298455
 */

import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import javax.imageio.ImageIO;

public class ColorImage {
    private BufferedImage image;
    private int width;
    private int height;
    private int depth;
    private int[][][] pixels; // [x/wdith][y/height][rgb]

    public ColorImage(String filename) {
        try 
        {
            File file = new File(filename);
            this.image = ImageIO.read(file);

            this.width = image.getWidth();
            this.height = image.getHeight();
            this.depth = image.getColorModel().getPixelSize();
            this.pixels = new int[this.width][this.height][3];
            for (int x = 0;x<width;x++)
            {
                for (int y = 0;y<height;y++)
                {
                    pixels[x][y] = getPixel(x, y);
                }
            }
        } 
        catch (IOException e) 
        {
            e.printStackTrace();
        }
    }

    public int getWidth() {
        return width;
    }

    public int getHeight() {
        return height;
    }

    public int getDepth() {
        return depth;
    }

    public int[][][] getPixels() {
        return pixels;
    } 

    // get the 3-channel value of a pixel at coord x,y
    // we assume the image is in 8-bit RGB color space
    public int[] getPixel(int x, int y) 
    {
        int mask = 0xFF; // least significant/rightmost 8-bits 
        int[] pixel = new int[3];
        int rgb = image.getRGB(x, y);
        pixel[0] = (rgb >> 16) & mask;  // red
        pixel[1] = (rgb >> 8) & mask;   // green
        pixel[2] = rgb & mask;          // blue
        return pixel;
    }

    // reduce the color space to a d-bit representation
    public void reduceColor(int d)
    {
        if (d>=8)
            return;
        int shift = 8-d;
        for (int x= 0;x<this.width;x++)
        {
            for (int y=0;y<this.height;y++)
            {
                int[] pixel = getPixel(x, y);

                // reducing color space to d-bit
                pixel[0] = pixel[0]>>shift; 
                pixel[1] = pixel[1]>>shift; 
                pixel[2] = pixel[2]>>shift; 

                // update our pixels array (that represents the image) to have 3-bit color space
                this.pixels[x][y][0] = pixel[0];
                this.pixels[x][y][1] = pixel[1];
                this.pixels[x][y][2] = pixel[2];
            }
        }
    }
}