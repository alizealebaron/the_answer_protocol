package malen.vue;

import java.awt.BorderLayout;

import javax.swing.JPanel;
import javax.swing.JScrollPane;
import javax.swing.border.Border;

/** Panel principale de l'application
 * @author  : Alizéa Lebaron
 * @author  : Tom Goureau
 * @author  : Trystan Baillobay
 * @version : 1.0.0 - 19/12/2024
 * @since   : 19/12/2024
 */



public class PanelPrincipal extends JPanel
{
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                 Attributs                                                                        */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	private PanelOutils      panelOutils;
	private PanelImage       panelImage;
	private Frame            framePrincipale;

	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/
	/*                                                                Controleurs                                                                       */
	/*--------------------------------------------------------------------------------------------------------------------------------------------------*/

	public PanelPrincipal (Frame frame, PanelImage panelI, PanelOutils panelO, JScrollPane scPanel)
	{
		// Initialisation des composants
		this.panelOutils = panelO;
		this.panelImage  = panelI;

		this.framePrincipale = frame;

		// Initialisation du panel
		this.setLayout(new BorderLayout());
		this.add(panelOutils, BorderLayout.NORTH );
		this.add(scPanel    , BorderLayout.CENTER);
	}
}
