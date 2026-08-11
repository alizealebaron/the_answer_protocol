package malen.vue;

import java.awt.FlowLayout;

import javax.swing.JColorChooser;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JSpinner;
import javax.swing.SpinnerNumberModel;

/** JColorChoose personnel
 * @author  : Alizéa Lebaron
 * @author  : Tom Goureau
 * @author  : Trystan Baillobay
 * @version : 1.0.0 - 19/12/2024
 * @since   : 19/12/2024
 */
 
public class ColorChooserPalette extends JColorChooser 
{
 
	private JSpinner colorDifferenceSpinner;
 
	public ColorChooserPalette(int value) 
	{
		super();
		 
		colorDifferenceSpinner = new JSpinner(new SpinnerNumberModel(value, 0, 255, 1));
		JLabel differenceLabel = new JLabel("Distance entre couleurs :");

		// Créer un panneau pour le spinner
        JPanel customPanel = new JPanel(new FlowLayout(FlowLayout.LEFT));
        customPanel.add(differenceLabel);
        customPanel.add(colorDifferenceSpinner);

        // Ajouter le panneau personnalisé à l'aperçu
        this.setPreviewPanel(customPanel);
	}

	public int getColorDifference() 
	{
        return (int) colorDifferenceSpinner.getValue();
    }
}
 
