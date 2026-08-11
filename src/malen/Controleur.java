package malen;

import java.awt.Color;
import java.awt.Graphics2D;
import java.awt.RenderingHints;
import java.awt.Shape;

import java.awt.Cursor;
import java.awt.Toolkit;
import javax.swing.SwingUtilities;

import malen.vue.FramePrincipale;
import malen.vue.FrameSecondaire;
import malen.modele.Couleur;
import malen.modele.Point;

import java.awt.image.BufferedImage;

/**
 * Classe Controleur
 * 
 * @author : Alizéa Lebaron
 * @author : Tom Goureau
 * @version : 1.0.0 - 16/12/2024
 * @date : 16/12/2024
 */

public class Controleur 
{
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                  Attributs                                                                       */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public static final String PIPETTE = "pipette";
	public static final String SOURIS = "souris";
	public static final String POT_DE_PEINTURE = "pot";
	public static final String SELECTION_RECTANGLE = "rectangle";
	public static final String SELECTION_OVALE = "cercle";
	public static final String EFFACE_FOND = "fond";
	public static final String TEXT = "text";
	public static final String LUMINOSITE = "lumi";
	public static final String CONTRASTE = "cont";

	private static final String REPERTOIRE = "../data/images/";
	private FramePrincipale   mainFrame;
	private FrameSecondaire   subFrame;


	private boolean onMainFrame;

	private String curseur;

	private Point point1;
	private Point point2;
	private BufferedImage subImage;

	private int densite = 30;

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                Controleurs                                                                       */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public Controleur() {
		mainFrame = new FramePrincipale(this); // Crée la fenêtre principale
		this.curseur = "souris";

		this.point1 = null;
		this.point2 = null;
		this.subImage = null;
		this.onMainFrame = true;
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                Accesseurs                                                                        */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public String        getCurseur      () { return this.curseur;     }
	public Point         getPoint1       () { return this.point1;      }
	public Point         getPoint2       () { return this.point2;      }
	public BufferedImage getSubImage     () { return this.subImage;    }
	public int           getDensite      () { return this.densite;     }

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                               Modificateurs                                                                      */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public void setPoint1  (Point point1       ) {this.point1  = point1;       }
	public void setPoint2  (Point point2       ) {this.point2  = point2;       }
	public void setDensite (int   densite      ) {this.densite = densite;      }

	public void setCurseur(String curseur) {
		this.curseur = curseur;
		Toolkit toolkit = Toolkit.getDefaultToolkit();
		switch (this.curseur) {
			case Controleur.SELECTION_RECTANGLE:
				this.mainFrame.setCursor(new Cursor(Cursor.CROSSHAIR_CURSOR));
				if (this.subFrame != null) {
					this.subFrame.setCursor(new Cursor(Cursor.CROSSHAIR_CURSOR));
				}
				break;

			case Controleur.SELECTION_OVALE:
				this.mainFrame.setCursor(new Cursor(Cursor.CROSSHAIR_CURSOR));
				if (this.subFrame != null) {
					this.subFrame.setCursor(new Cursor(Cursor.CROSSHAIR_CURSOR));
				}
				break;

			case Controleur.PIPETTE:
				this.mainFrame.setCursor(toolkit.createCustomCursor(
						toolkit.getImage(REPERTOIRE + "pipette-curseur.png"), new java.awt.Point(0, 25), "Pipette"));
				if (this.subFrame != null) {
					this.subFrame.setCursor(
							toolkit.createCustomCursor(toolkit.getImage(REPERTOIRE + "pipette-curseur.png"),
									new java.awt.Point(0, 25), "Pipette"));
				}
				break;

			case Controleur.POT_DE_PEINTURE:
				this.mainFrame.setCursor(toolkit.createCustomCursor(toolkit.getImage(REPERTOIRE + "peinture.png"),
						new java.awt.Point(23,23), "Pot"));
				if (this.subFrame != null) {
					this.subFrame.setCursor(toolkit.createCustomCursor(toolkit.getImage(REPERTOIRE + "peinture.png"),
							new java.awt.Point(23,23), "Pot"));
				}
				break;

			case Controleur.EFFACE_FOND:
				this.mainFrame.setCursor(toolkit.createCustomCursor(toolkit.getImage(REPERTOIRE + "goutte.png"),
						new java.awt.Point(16, 16), "Transparence"));
				if (this.subFrame != null) {
					this.subFrame
							.setCursor(toolkit.createCustomCursor(toolkit.getImage(REPERTOIRE + "goutte.png"),
									new java.awt.Point(16, 16), "Transparence"));
				}
				break;

			case Controleur.TEXT:
				this.mainFrame.setCursor(new Cursor(Cursor.TEXT_CURSOR));
				if (this.subFrame != null) {
					this.subFrame.setCursor(new Cursor(Cursor.TEXT_CURSOR));
				}
				break;

			default:
				this.mainFrame.setCursor(new Cursor(Cursor.DEFAULT_CURSOR));
				if (this.subFrame != null) {
					this.subFrame.setCursor(new Cursor(Cursor.DEFAULT_CURSOR));
				}
				break;
		}
		if (!curseur.equals(SELECTION_RECTANGLE)) {
			resetSelection(); // Réinitialiser la sélection quand on sort du mode sélection rectangle
		}
	}


	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                 Méthodes                                                                         */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	// Méthode pour démarrer l'application
	public void startApplication() {
		SwingUtilities.invokeLater(() -> {
			mainFrame.setVisible(true);
		});
	}


	public void nouvelleFenetre() {
		subFrame = new FrameSecondaire(this.mainFrame, this);
		subFrame.setVisible(true);
		mainFrame.actualiserMenu(false);
	}

	public void resetSelection() {
		point1 = null;
		point2 = null;
		subImage = null;
	}

	// Méthode pour obtenir la couleur actuelle


	public void setSubImage(BufferedImage image) {
		this.subImage = image;
	}
	
	public void setOnMainFrame() {
		this.onMainFrame = true;
	}

	public void setOnSecondFrame() {
		this.onMainFrame = false;
	}

	public boolean isOnMainFrame() {
		return this.onMainFrame;
	}

	public boolean isOnSecondFrame() {
		return !this.onMainFrame;
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*---------------------------------------------------------Action avec le click---------------------------------------------------------------------*/
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	/**
	 * Méthode permettant de faire une action pendant un click selon l'état du
	 * curseur
	 * 
	 */
	public void onClickLeft(BufferedImage biImage, Color coulPixel, int x, int y) 
	{
		switch (this.curseur) {
			case Controleur.SOURIS:
				break;

			case Controleur.PIPETTE:
				if (this.isOnMainFrame()) {
					this.mainFrame.setColor(coulPixel);
				}
				if (this.isOnSecondFrame()) {
					this.subFrame.setColor(coulPixel);
				}
				break;

			case Controleur.POT_DE_PEINTURE:
			
				if (this.isOnMainFrame()) 
				{
					fill(biImage, coulPixel, this.mainFrame.getCurrentColor(), x, y, this.densite);
				}
				if (this.isOnSecondFrame()) 
				{
					fill(biImage, coulPixel, this.subFrame.getCurrentColor(), x, y, this.densite);
				}
				
				break;

			case Controleur.EFFACE_FOND:
				fondTransparent(biImage, coulPixel, x, y, this.densite);

				break;

			case Controleur.SELECTION_RECTANGLE:
				if (this.subImage != null) { // peut poser des problemes, mettre verif sur point1 et point2
					if (this.isOnMainFrame()) {
						this.mainFrame.pasteSubImage();
					} else {
						this.subFrame.pasteSubImage();
					}
					this.setCurseur(Controleur.SOURIS);
					this.resetSelection();
				}
				break;

			case Controleur.SELECTION_OVALE:
				if (this.subImage != null) { // peut poser des problemes, mettre verif sur point1 et point2
					if (this.isOnMainFrame()) {
						this.mainFrame.pasteSubImage();
					} else {
						this.subFrame.pasteSubImage();
					}
					this.setCurseur(Controleur.SOURIS);
					this.resetSelection();
				}
				break;

			default:
				break;

		}
	}

	public void onClickRight()
	{
		if (this.subFrame != null)
		{
			if (onMainFrame)
			{
				this.subFrame.repaint();
			}
			else
			{
				this.mainFrame.repaint();
			}
		}
	}

	public void createRectangleSubImage(BufferedImage image) {
		if (image == null) {
			this.subImage = null;
		} else {
			if (point1 != null && point2 != null && point1.x() != point2.x() && point1.y() != point2.y()) {
				int x1 = Math.min(point1.x(), point2.x());
				int y1 = Math.min(point1.y(), point2.y());
				int width = Math.abs(point1.x() - point2.x());
				int height = Math.abs(point1.y() - point2.y());
				subImage = image.getSubimage(x1, y1, width, height);
			}
		}
	}

	public void createOvalSubImage(BufferedImage image) {

		if (image == null) {
			this.subImage = null;
		} else {
			if (point1 != null && point2 != null && point1.x() != point2.x() && point1.y() != point2.y()) {
				int x1 = Math.min(point1.x(), point2.x());
				int y1 = Math.min(point1.y(), point2.y());
				int width = Math.abs(point1.x() - point2.x());
				int height = Math.abs(point1.y() - point2.y());

				// Crée une image vide avec un canal alpha (transparence)
				BufferedImage ovalImage = new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);

				// Crée un Graphics2D pour dessiner sur l'image ovale
				Graphics2D g2d = ovalImage.createGraphics();

				// Activer l'anti-aliasing pour des bords plus lisses
				g2d.setRenderingHint(RenderingHints.KEY_ANTIALIASING, RenderingHints.VALUE_ANTIALIAS_ON);
				g2d.setRenderingHint(RenderingHints.KEY_RENDERING, RenderingHints.VALUE_RENDER_QUALITY);

				// Créer une forme ovale
				Shape oval = new java.awt.geom.Ellipse2D.Double(0, 0, width, height);

				// Dessiner l'image source à l'intérieur de l'ovale
				g2d.setClip(oval); // Limite le dessin à la région ovale
				g2d.drawImage(image, -x1, -y1, null); // Ajuste pour positionner la bonne partie de l'image source

				g2d.dispose();
				this.subImage = ovalImage;
			}
		}
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*---------------------------------------------------------Liaison modele couleur-------------------------------------------------------------------*/
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public BufferedImage fill(BufferedImage biOriginal, Color colOld, Color colNew, int x, int y, int distance) {
		return Couleur.fill(biOriginal, colOld, colNew, x, y, this.densite);
	}

	public BufferedImage fondTransparent(BufferedImage biOriginal, Color colOld, int x, int y, int distance) {
		return Couleur.fondTransparent(biOriginal, colOld, x, y, this.densite);
	}

	public BufferedImage changerContraste(BufferedImage biOriginal, int contraste) {
		return Couleur.changerContraste(biOriginal, contraste);
	}

	public BufferedImage changerLuminosite(BufferedImage biOriginal, int luminosite) {
		return Couleur.changerLuminosite(biOriginal, luminosite);
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*------------------------------------------------------------------main----------------------------------------------------------------------------*/
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public static void main(String[] args) {
		// Lancer l'application depuis le contrôleur
		Controleur controleur = new Controleur();
		controleur.startApplication();
	}
}
