package lang

import "github.com/LordPax/aicli/utils"

var FR_STRINGS = LangString{
	"usage":                   "CLI pour utiliser des modèles d'IA",
	"output-desc":             "Répertoire de sortie",
	"output-dir-empty":        "Le répertoire de sortie est vide",
	"silent":                  "Désactiver l'impression du journal sur stdout",
	"no-args":                 "Aucun argument fourni",
	"no-command":              "Aucune commande fournie",
	"unknown-sdk":             "Sdk inconnu \"%s\"",
	"sdk-model-usage":         "Sélectionner un modèle",
	"text-usage":              "Générer du texte à partir d'un prompt",
	"text-sdk-usage":          "Sélectionner un sdk",
	"text-temp-usage":         "Définir la température",
	"text-system-usage":       "Instruction avec rôle système (utilisez \"-\" pour stdin)",
	"text-history-usage":      "Sélectionner un historique",
	"text-clear-usage":        "Effacer l'historique",
	"text-file-usage":         "Fichier texte à utiliser",
	"text-input":              "(\"exit\" pour quitter) " + utils.Blue + "user> " + utils.Reset,
	"text-list-history-usage": "Lister l'historique",
	"type-required":           "Le type est requis",
	"apiKey-required":         "La clé API est requise",
	"empty-file":              "Le fichier est vide",
	"empty-history":           "L'historique \"%s\" est vide\n",
}
