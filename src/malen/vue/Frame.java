package malen.vue;

import java.awt.BorderLayout;
import java.awt.Color;
import java.awt.Cursor;
import java.awt.Dimension;
import java.awt.Graphics2D;
import java.awt.event.MouseEvent;
import java.awt.image.BufferedImage;

import java.io.File;

import javax.swing.*;
import javax.swing.filechooser.FileNameExtensionFilter;

import malen.Controleur;
import malen.modele.Point;

/** Frame parente
 * @author  : Alizéa Lebaron
 * @author  : Tom Goureau
 * @author  : Trystan Baillobay
 * @version : 1.0.0 - 19/12/2024
 * @since   : 19/12/2024
 */

public abstract class Frame extends JFrame 
{
	protected Controleur controleur;
	protected JScrollPane      scPanelPrincipal;

	protected PanelPrincipal   panelPrincipal;
	protected PanelOutils      panelOutils;
	protected PanelImage       panelImage;

	protected Color currentColor = Color.BLACK; // La couleur actuelle, par défaut noire

	public Frame(Controleur controleur) {
		this.controleur = controleur;

		setSize(1650, 1080);
		this.setExtendedState(JFrame.MAXIMIZED_BOTH);
		setLayout(new BorderLayout());

		// Initialisation du panel outil et du panel image
		this.panelImage  = new PanelImage (this);
		this.panelOutils = new PanelOutils(this);

		this.scPanelPrincipal = new JScrollPane(this.panelImage); // Envelopper l'image dans un JScrollPane
		this.scPanelPrincipal.setPreferredSize(new Dimension(800, 600)); // Taille du panneau d'affichage de l'image

		// Ajouter le panel principale
		this.panelPrincipal   = new PanelPrincipal (this, this.panelImage, this.panelOutils, this.scPanelPrincipal);
		this.add(this.panelPrincipal, BorderLayout.CENTER);

		try 
		{
			// Définit le look-and-feel natif du système
			UIManager.setLookAndFeel(UIManager.getSystemLookAndFeelClassName());
		} 
		catch (Exception e) 
		{
			System.err.println("Impossible d'appliquer le look-and-feel natif : " + e.getMessage());
		}

		// Afficher la fenêtre
		setLocationRelativeTo(null);
		setVisible(true);
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                Accesseurs                                                                        */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public Color         getCurrentColor ( ) { return this.currentColor                  ; }
	public Point         getPoint1       ( ) { return this.controleur.getPoint1       ( ); }
	public Point         getPoint2       ( ) { return this.controleur.getPoint2       ( ); }
	public BufferedImage getSubImage     ( ) { return this.controleur.getSubImage     ( ); }

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                               Modificateurs                                                                      */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public void setPoint1   ( Point         point1 ) { this.controleur.setPoint1   ( point1 );}
	public void setPoint2   ( Point         point2 ) { this.controleur.setPoint2   ( point2 );}
	public void setSubimage ( BufferedImage image  ) { this.controleur.setSubImage ( image  );}

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                 Méthodes                                                                         */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	// Méthode pour ouvrir le dialogue d'importation d'image
	public void importImage() 
	{
		JFileChooser fileChooser = new JFileChooser();
        fileChooser.setDialogTitle("Sélectionnez une image");

        // Ajouter un filtre pour les fichiers GIF et PNG
        FileNameExtensionFilter filter = new FileNameExtensionFilter("Images (GIF, PNG)", "gif", "png");
        fileChooser.setFileFilter(filter);

        // Empêcher l'utilisateur de sélectionner des répertoires
        fileChooser.setFileSelectionMode(JFileChooser.FILES_ONLY);

        // Afficher le dialogue
        int result = fileChooser.showOpenDialog(this);
        if (result == JFileChooser.APPROVE_OPTION) 
		{
            File selectedFile = fileChooser.getSelectedFile();
			// Importer l'image dans le panneau
			this.panelImage.importImage(selectedFile.getAbsolutePath());

			// Mettre à jour la taille du JScrollPane selon la taille de l'image
			Dimension imageSize = this.panelImage.getImageSize();
			this.panelImage.setPreferredSize(imageSize); // Mettre à jour la taille du panel

			this.scPanelPrincipal.revalidate(); // Revalider le JScrollPane pour appliquer les nouvelles dimensions
			this.scPanelPrincipal.repaint   (); // Redessiner le JScrollPane
        }
	}

	/* ------------------------------------------------------------ */
	/*                   Passerelle Outils-Image                    */
	/* ------------------------------------------------------------ */

	/*                  Accesseurs                  */
	/* -------------------------------------------- */
	
	public void          repaintImage   ( ) {        this.panelImage.repaint        ( ) ; }
	public boolean       isImage        ( ) { return this.panelImage.isImage        ( ) ; }
	public BufferedImage getImage       ( ) { return this.panelImage.getImage       ( ) ; }
	public Color         getColorText   ( ) { return this.panelImage.getColorText   ( ) ; }
	public Double        getRotateAngle ( ) { return this.panelImage.getRotateAngle ( ) ; }

	public char          getOutil       ( ) { return this.panelOutils.getOutil      ( ) ; }

	/*                 Modificateurs                */
	/* -------------------------------------------- */

	public void setRotateAngle ( double angle ) { this.panelImage.setRotateAngle (angle); }
	public void setFontImage   ( File   image ) { this.panelImage.setFontImage   (image); }

	/*                    Autres                    */
	/* -------------------------------------------- */

	public void updateTextFont (JComboBox<String> fontBox, JSpinner sizeSpinner, JCheckBox boldCheck, JCheckBox italicCheck) { this.panelImage.updateTextFont(fontBox, sizeSpinner, boldCheck, italicCheck);}

	public void updateTextColor ( Color c ) 
	{ 
		this.panelImage.setTextColor(c);
		this.panelImage.getTextField().setForeground(c);
	}

	/**
	 * Permet d'afficher le panel de modification du texte
	 * 
	 */
	public void afficherPanelText() { this.panelOutils.afficherPanelText(); }	

	public void showOutilSlider(char outil) { this.panelOutils.showOutilSlider(outil);}

	/* ------------------------------------------------------------ */
	/*                     Gestion des outils                       */
	/* ------------------------------------------------------------ */

	public void changerContraste  (int  value) { this.controleur.changerContraste  (this.panelImage.getImage(), value); }
	public void changerLuminosite (int  value) { this.controleur.changerLuminosite (this.panelImage.getImage(), value); }

	public void afficherSlider    (char outil) { this.panelOutils.showOutilSlider   (outil)                           ; }

	/* ------------------------------------------------------------ */
	/*                     Liaison Controleur                       */
	/* ------------------------------------------------------------ */

	public void setCurrentColor (Color c) { this.setColor(c);}

	public void switchCurseur(String curseur)
	{
		if (this.controleur.getCurseur().equals(curseur))
		{
			this.controleur.setCurseur(Controleur.SOURIS);
		}
		else
		{
			this.controleur.setCurseur(curseur);
		}
	}

	public void setCursor(Cursor nouveauCurseur)
	{
		this.scPanelPrincipal.setCursor(nouveauCurseur);
	}

	public boolean isCurseurOn(String curseur)
	{
		return this.controleur.getCurseur().equals(curseur);
	}

	public int  getDensite (     ) { return this.controleur.getDensite( );}
	public void setDensite (int d) {        this.controleur.setDensite(d);}

	public void chooseColor()
	{
		// Créer une instance de ColorChooserPalette
		ColorChooserPalette colorChooser = new ColorChooserPalette(this.getDensite());

		// Afficher le sélecteur de couleurs
		int option = JOptionPane.showConfirmDialog
		(
			null,
			colorChooser,
			"Choisir une couleur",
			JOptionPane.OK_CANCEL_OPTION
		);
	
		if (option == JOptionPane.OK_OPTION) 
		{
			// Récupérer la couleur et la différence
			Color selectedColor = colorChooser.getColor();
			int colorDifference = colorChooser.getColorDifference();
			
			if (selectedColor != null)  
			{
				this.setColor(selectedColor);
				this.setDensite(colorDifference);
			}
		}
	}

	public void setColor (Color selectedColor) 
	{
		this.currentColor = selectedColor;
		this.updateButton();
	}

	public abstract void updateButton();

	public void saveImage() {}

	public void saveSousImage () {}

	public void saveImage(String fileName) {}

	public void switchRetournementHorizontal()
	{
		panelImage.switchFlipHorizontal();
	}

	public void switchRetournementVertical()
	{
		panelImage.switchFlipVertical();
	}

	public void changerContraste(BufferedImage bi, int value)
	{
		this.controleur.changerContraste(bi, value);
	}

	public void changerLuminosite(BufferedImage bi, int value)
	{
		this.controleur.changerLuminosite(bi, value);
	}

	public void onClickLeft(BufferedImage biImage, int x, int y, Color coulPixel)
	{
		this.controleur.onClickLeft(biImage, coulPixel, x, y);
	}

	public abstract void onClickRight(MouseEvent e);

	public void createSubImage(BufferedImage image)
	{
		if (this.isCurseurOn(Controleur.SELECTION_RECTANGLE))
		{
			this.controleur.createRectangleSubImage(image);
		}
		if (this.isCurseurOn(Controleur.SELECTION_OVALE))
		{
			this.controleur.createOvalSubImage(image);
		}
	}

	public BufferedImage pasteSubImage()
	{
		BufferedImage image = this.panelImage.getImage();

		if (this.getSubImage() != null && image != null && this.getPoint1() != null
				&& this.getPoint2() != null)
		{
			// Déterminer les coordonnées où coller la subimage
			int x = Math.min(this.getPoint1().x(), this.getPoint2().x());
			int y = Math.min(this.getPoint1().y(), this.getPoint2().y());

			// Dessiner la subimage sur l'image principale
			Graphics2D g2d = image.createGraphics();
			g2d.drawImage(this.getSubImage(), x, y, null);
			g2d.dispose();

			// Réinitialiser la sélection et la subimage
			this.setPoint1(null);
			this.setPoint2(null); // Réinitialiser le deuxième point
			this.createSubImage(null);
		}
		return image;
	}

	public void nouvelleFenetre() {}

	public void export() {}

	public void setOnMainFrame()
	{
		this.controleur.setOnMainFrame();
	}

	public void setOnSecondFrame()
	{
		this.controleur.setOnSecondFrame();
	}

	public boolean isOnMainFrame()
	{
		return this.controleur.isOnMainFrame();
	}

	public boolean isOnSecondFrame()
	{
		return this.controleur.isOnSecondFrame();
	}

	public abstract boolean isMainFrame();
}
