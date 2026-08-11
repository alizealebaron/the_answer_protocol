package malen.modele;

/** Classe Controleur
  * @author : Alizéa Lebaron
  * @version : 1.0.0 - 16/12/2024
  * @date : 16/12/2024
  */

import java.awt.image.BufferedImage;
import java.util.LinkedList;
import java.util.Queue;

import java.awt.Color;

public class Couleur 
{
	/* ------------------------------------------------------------------------------------------------------------------------------------------------ */
	/*                                                                Pour le controller                                                                */
	/* ------------------------------------------------------------------------------------------------------------------------------------------------ */

	/** Méthode de remplissage
	 * @param biOriginal L'image sans la modification
	 * @param colOld La couleur a remplacé par la nouvelle
	 * @param colNew La couleur qui remplacera l'ancienne
	 * @param x Coordonnées x du point initial
	 * @param y Coordonnées y du point initial
	 * @return L'image modifiée
	 */
	public static BufferedImage fill(BufferedImage biOriginal, Color colOld, Color colNew, int x, int y, int distance)
	{
		// Initialisation de la liste des points
		Queue<Point> file = new LinkedList<Point>();
		file.add ( new Point ( x, y ) );

		// Récupération des couleurs en int
		int intCoulOld = colOld.getRGB() & 0xFFFFFF;
		int intCoulNew = colNew.getRGB() & 0xFFFFFF;

		if (intCoulOld == intCoulNew) { return biOriginal; }

		// Changement de tous les points présents dans la file
		while ( ! file.isEmpty() )
		{
			// On retire le point de la liste
			Point p = file.remove();

			// On vérifie que le point est dans l'image et que la couleur est ressemblante
			if ( p.x() >= 0 && p.x() < biOriginal.getWidth  () && p.y() >= 0 && p.y() < biOriginal.getHeight () && distance (intCoulOld, biOriginal.getRGB(p.x(), p.y() ) & 0xFFFFFF ) < distance)
			{
				// On change la couleur
				biOriginal.setRGB ( p.x(), p.y(), colNew.getRGB());

				// On ajoute tous les voisins à la file
				file.add ( new Point ( p.x()+1, p.y()   ) );
				file.add ( new Point ( p.x()-1, p.y()   ) );
				file.add ( new Point ( p.x()  , p.y()-1 ) );
				file.add ( new Point ( p.x()  , p.y()+1 ) );
			}
		}

		return biOriginal;
	}

	/** Méthode permettant de rendre le fond transparent
	 * @param biOriginal Image original à modifier
	 * @param colOld Couleur du fond à rendre transparent
	 * @param x Coordonnées x du point de référence
	 * @param y Coordonnées y du point de référence
	 * @return L'image modifiée
	 */
	public static BufferedImage fondTransparent(BufferedImage biOriginal, Color colOld, int x, int y, int distance)
	{
		// Initialisation de la liste des points
		Queue<Point> file = new LinkedList<Point>();
		file.add ( new Point ( x, y ) );

		// Récupération des couleurs en int
		int intCoulOld  = colOld.getRGB() & 0xFFFFFF;
		int coulTranspa = 0x00000000; // Couleur complètement transparente

		if (intCoulOld == coulTranspa) { return biOriginal; }

		// Changement de tous les points présents dans la file
		while ( ! file.isEmpty() )
		{
			// On retire le point de la liste
			Point p = file.remove();

			// On vérifie que le point est dans l'image et que la couleur est ressemblante
			if ( p.x() >= 0 && p.x() < biOriginal.getWidth  () && p.y() >= 0 && p.y() < biOriginal.getHeight () && distance (intCoulOld, biOriginal.getRGB(p.x(), p.y() ) & 0xFFFFFF ) < distance)
			{
				// On change la couleur
				biOriginal.setRGB ( p.x(), p.y(), coulTranspa);

				// On ajoute tous les voisins à la file
				file.add ( new Point ( p.x()+1, p.y()   ) );
				file.add ( new Point ( p.x()-1, p.y()   ) );
				file.add ( new Point ( p.x()  , p.y()-1 ) );
				file.add ( new Point ( p.x()  , p.y()+1 ) );
			}
		}

		return biOriginal;
	}

	/** Permet de modifier le contraste d'une image
	 * @param biOriginal L'image a modifiée
	 * @param contraste Le constraste à ajouter
	 * @return La nouvelle image
	 */
	public static BufferedImage changerContraste(BufferedImage biOriginal, int contraste)
	{
		Color couleur;

		for (int x = 0; x < biOriginal.getWidth(); x++)
		{
			for (int y = 0; y < biOriginal.getHeight(); y++)
			{
				couleur = new Color (biOriginal.getRGB(x, y) & 0xFFFFFF);

				int rouge = verifBorneCouleur((int) (couleur.getRed  () + contraste / 100.0 * (couleur.getRed  ()-127)));
				int vert  = verifBorneCouleur((int) (couleur.getGreen() + contraste / 100.0 * (couleur.getGreen()-127)));
				int bleu  = verifBorneCouleur((int) (couleur.getBlue () + contraste / 100.0 * (couleur.getBlue ()-127)));

				biOriginal.setRGB(x, y, (new Color (rouge, vert, bleu).getRGB()));
			}
		}

		return biOriginal;
	}

	public static BufferedImage changerLuminosite(BufferedImage biOriginal, int luminosite)
	{
		Color couleur;

		for (int x = 0; x < biOriginal.getWidth(); x++)
		{
			for (int y = 0; y < biOriginal.getHeight(); y++)
			{
				couleur = new Color (biOriginal.getRGB(x, y) & 0xFFFFFF);

				int rouge = verifBorneCouleur(couleur.getRed()   + luminosite);
				int vert  = verifBorneCouleur(couleur.getGreen() + luminosite);
				int bleu  = verifBorneCouleur(couleur.getBlue()  + luminosite);

				biOriginal.setRGB(x, y, (new Color (rouge, vert, bleu).getRGB()));
			}
		}

		return biOriginal;
	}

	/* ------------------------------------------------------------------------------------------------------------------------------------------------ */
	/*                                                                  Pour Couleur                                                                    */
	/* ------------------------------------------------------------------------------------------------------------------------------------------------ */

	/** Permet de comparer la distance entre deux couleurs
	 * @param coul1 Première couleur à comparée
	 * @param coul2 Deuxième couleur à comparée
	 * @return La distance séparant les deux couleurs
	 */
	private static double distance (int coul1, int coul2)
	{
		Color color1 = new Color (coul1);
		Color color2 = new Color (coul2);

		return Math.sqrt(Math.pow((color1.getRed() - color2.getRed()), 2 ) + Math.pow((color1.getGreen() - color2.getGreen()), 2 ) + Math.pow((color1.getBlue() - color2.getBlue()), 2 ));
	}


	/** Permet de vérifier que la couleur est dans les bornes
	 * @param coul La couleur a vérifier
	 * @return La couleur changée si elle dépasse les bornes
	 */
	private static int verifBorneCouleur (int coul)
	{
		// Si la couleur est trop grande ou trop petite, on la ramène à la borne inférieur ou supérieur
		if (coul > 255) { coul = 255;} 
		if (coul <   0) { coul =   0;}

		return coul;
	}

	/** Méthode pour tester sans le contrôleur
	 * @param args Arguments donnés
	 */
	// public static void main(String[] args) 
	// {
	// 	BufferedImage biOriginal = null;

	// 	try 
	// 	{
	// 		biOriginal = ImageIO.read(new File("wooper.png"));
	// 	} 
	// 	catch (Exception e) 
	// 	{
	// 		e.printStackTrace();
	// 	}

	// 	Color coulOrig = new Color(248,248,248);
	// 	Color coulNew  = new Color(  0,255,  0);

	// 	//fill(biOriginal, coulOrig, coulNew, 1, 1);
	// 	//fondTransparent(biOriginal, coulOrig, 1, 1);
	// 	//changerContraste(biOriginal, -60);
	// 	changerLuminosite(biOriginal, -110);

	// 	try 
	// 	{
	// 		ImageIO.write(biOriginal, "png", new File("Arceus6.png"));
	// 	} 
	// 	catch (Exception e) 
	// 	{
	// 		e.printStackTrace();
	// 	}
	// }
	
}
