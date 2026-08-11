package malen.modele;

import java.awt.*;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;
import javax.swing.JOptionPane;

public class Rotation
{
    private double  angle;
    private boolean flipHorizontal;
    private boolean flipVertical;

    public Rotation(double angle, boolean flipHorizontal, boolean flipVertical)
	{
        this.angle          = angle;
        this.flipHorizontal = flipHorizontal;
        this.flipVertical   = flipVertical;
    }

    public BufferedImage applyTransformations(BufferedImage image)
	{
		BufferedImage outputImage = image;
		if (image != null)
		{
			int originalWidth  = image.getWidth();
			int originalHeight = image.getHeight();
		
			double radians = Math.toRadians(angle);
			int newWidth   = (int) Math.round(Math.abs(originalWidth  * Math.cos(radians)) + 
											Math.abs(originalHeight * Math.sin(radians)));
			int newHeight  = (int) Math.round(Math.abs(originalWidth  * Math.sin(radians)) + 
											Math.abs(originalHeight * Math.cos(radians)));
		
			outputImage = new BufferedImage(newWidth, newHeight, BufferedImage.TYPE_INT_ARGB);
			Graphics2D g2d = outputImage.createGraphics();
		
			// Activer l'anti-aliasing pour une meilleure qualité et ne sauvegarder que l'image
			g2d.setRenderingHint(RenderingHints.KEY_INTERPOLATION, RenderingHints.VALUE_INTERPOLATION_BILINEAR);
			g2d.setRenderingHint(RenderingHints.KEY_RENDERING, RenderingHints.VALUE_RENDER_QUALITY);
			g2d.setRenderingHint(RenderingHints.KEY_ANTIALIASING, RenderingHints.VALUE_ANTIALIAS_ON);
		
			g2d.translate(newWidth / 2, newHeight / 2);
		
			if (flipHorizontal) { g2d.scale(-1, 1); }
			if (flipVertical)   { g2d.scale(1, -1); }
		
			g2d.rotate(radians);
		
			g2d.translate(-originalWidth / 2, -originalHeight / 2);
		
			g2d.drawImage(image, 0, 0, null);
			g2d.dispose();
		}

		return outputImage;
	}

	public void saveImageToFile(BufferedImage image, String filePath)
	{
        BufferedImage transformedImage = applyTransformations(image);

        try
		{
            File outputFile = new File(filePath);
            ImageIO.write(transformedImage, "png", outputFile);
            JOptionPane.showMessageDialog(null, "Image sauvegardée avec succès.", "Succès", JOptionPane.INFORMATION_MESSAGE);
        } catch (IOException e)
		{
            e.printStackTrace();
            JOptionPane.showMessageDialog(null, "Erreur lors de la sauvegarde.", "Erreur", JOptionPane.ERROR_MESSAGE);
        }
    }

	public BufferedImage getImage(BufferedImage image)
	{
        BufferedImage transformedImage = applyTransformations(image);
		return transformedImage;
    }
}