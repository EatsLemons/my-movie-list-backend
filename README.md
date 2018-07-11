# My movie list backend API

## Requests
### Autocomplete
/api/v1/autocomplete/{query}

```json
{
	"autocomplete_movie": [
		{
			"name": "Перевал",
			"actors": [
				"Vasiliy Livanov",
				"Aleksandr Kaydanovskiy",
				"Aleksandr Pashutin"
			],
			"release_year": "1988",
			"image_url": "https://image.tmdb.org/t/p/w500/mc1FP17S1a5r8ZKFTEDFojI55wt.jpg"
		},
		{
			"name": "Высокий перевал",
			"actors": [
				"Natalia Naum",
				"Konstantin Stepankov",
				"Lyubov Bogdan"
			],
			"release_year": "1981",
			"image_url": "https://image.tmdb.org/t/p/w500/7OYsJmwqvH39SaGpD0vhRGinfVF.jpg"
		}
	],
	"system_info": {}
}
```
