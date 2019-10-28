## INSTALL

```
brew cask install wkhtmltopdf
go get github.com/thejini3/blog-to-pdf
cd $GOPATH/src/github.com/thejini3/blog-to-pdf
go install
```


## USAGE
```
Usage of blog-to-pdf:

  $ blog-to-pdf init <dir_name>
  
  # Edit auto-generated .ini file, then -

  $ blog-to-pdf <config_file.ini>

```

## File/Directory structure
```
/<your_project>/config.ini
/<your_project>/custom.css
/<your_project>/string_replaces.json
/<your_project>/urls.txt
/<your_project>/original-html/
/<your_project>/combined-html/
/<your_project>/pdf/
```

## `config.ini` sample

To generate pdf after modifying your configs, use `generate_pdf = true`

It's disabled by default, since every blog needs some modification first

```
protocol = https://

domain = your_blog.com

sitemap_url = your_blog.com/sitemap.xml

article_per_pdf = 10

# This is one of the important value, since we will merge (article_per_pdf) 10 article in a single PDF
# so, which portion of the HTML will be merged in the main Layout? Ex: 'div#content', 'div.post', 'article'
# for "id", use "$" instead of "#"
article_parent_element = $content

# for "id", use "$" instead of "#"
article_title_class = .post h2$entry-title

# for "id", use "$" instead of "#"
elements_to_remove = footer, aside, .respond, $wpcom-block-editor-styles-css, link[rel="dns-prefetch"], .wpcnt, $jp-post-flair, .post-nav-wrapper, .menu-wrapper

# There will be in need of some REPLACES, that's why had to use JSON file,
# Make sure that's valid JSON file

string_replaces_file = string_replaces.json

# Force Re-fetch htmls from server
force_html_fetch = false

# Force Re-fetch urls from sitemap / by wget
force_urls_fetch = true

# -1 => fetch all url
# otherwise, you can limit number of urls
limit_urls = -1

# Generate pdf or not, if false then only combined-html files will be created!
generate_pdf = false

# A0        =>	841 x 1189 mm
# A1        =>	594 x 841 mm
# A2        =>	420 x 594 mm
# A3        =>	297 x 420 mm
# A4        =>	210 x 297 mm, 8.26
# A5        =>	148 x 210 mm
# A6        =>	105 x 148 mm
# A7        =>	74 x 105 mm
# A8        =>	52 x 74 mm
# A9        =>	37 x 52 mm
# B0        =>	1000 x 1414 mm
# B1        =>	707 x 1000 mm
# B10       =>	31 x 44 mm
# B2        =>	500 x 707 mm
# B3        =>	353 x 500 mm
# B4        =>	250 x 353 mm
# B5        =>	176 x 250 mm, 6.93
# B6        =>	125 x 176 mm
# B7        =>	88 x 125 mm
# B8        =>	62 x 88 mm
# B9        =>	33 x 62 mm
# C5E       =>	163 x 229 mm
# Comm10E   =>	105 x 241 mm, U.S. Common 10 Envelope
# Custom    =>	Unknown, or a user defined size.
# DLE       =>	110 x 220 mm
# Executive =>	7.5 x 10 inches, 190.5 x 254 mm
# Folio     =>	210 x 330 mm
# Ledger    =>	431.8 x 279.4 mm
# Legal     =>	8.5 x 14 inches, 215.9 x 355.6 mm
# Letter    =>	8.5 x 11 inches, 215.9 x 279.4 mm
# Tabloid   =>	279.4 x 431.8 mm

pdf_size = A7


# "Landscape" or "Portrait"
pdf_orientation = Portrait


# UI
custom_css_file = custom.css

# Margin / White spaces for pdf (mm)
pdf_margin_top = 3
pdf_margin_left = 1
pdf_margin_right = 1
pdf_margin_bottom = 3
```


## EXAMPLE

```
$ blog-to-pdf init amarspondon
$ cd amarspondon
$ blog-to-pdf config.ini
```

