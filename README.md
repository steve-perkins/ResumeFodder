Resume Template Processor
=========================

Intro
-----
A command-line utility for generating resumes or CV's, inspired by
[HackMyResume](https://github.com/hacksalot/HackMyResume).

The processor uses resume data files in [JSON Resume](https://github.com/jsonresume/resume-schema)
format.  Your resume data is applied to a template, producing a polished Microsoft Word file as
output.  The separation between the resume data and the formatting allows you to quickly swap out
your resume's style with no painful cut-n-pasting.

Also, this makes it simple for you to place your resume data in source control.  Track and diff
changes to its substantive content over time, apart from formatting changes.  Create multiple
branches, for custom-tailored versions of your resume that have different buzzwords emphasized
depending on the target audience.

Differences between this and HackMyResume
-----------------------------------------

### The Good

* Written in Go, this utility compiles to a small **self-contained executable**.  There's no need to
  install Node.js or manage NPM on your machine.

* This utility supports some **optional extensions to the standard JSON Resume schema**.  These extra
  fields do not break compatibility... they would simply be ignored if you used your data file with
  HackMyResume or another JSON Resume processor.  However, you would of course need to manually
  reorganize this extra data into standard fields if you want other processors to use them.

    * Top-level `highlights`.  This allows you to include some bullet points in addition to (or
      instead of) a normal `summary` paragraph.

    * An `additionalWork` section, aside from the regular `work` entries.  This allows templates
      to apply different formatting to a certain portion of your work history.  For example, if you
      have a long work history then you might want to show full details for recent jobs, and
      abbreviated listings for older positions.  Likewise if you have internships or volunteer
      work that you would like to include, yet keep separate from your primary job history.

    * `workLabel` and `additionalWorkLabel` fields, to tell templates how to present the `work`
      and `additionalWork` sections when both are used (e.g. "Recent Experience" versus "Prior
      Experience (further detail available upon request)").

    * An `additionalPublications` section, aside from the regular `publications` entries.  This
      allows certain publications to receive different treatment.  For example, if you were a
      technical reviewer on someone else's book, then you might want to present that separately
      from books in which you were the actual author or co-author.

    * `publicationsLabel` and `additionalPublicationsLabel` fieldds, to tell templates how to present
      the `publications` and `additionalPublications` sections when both are used (e.g.
      "Publications (author)" versus "Additional Publications (technical reviewer)").

    * An `ISBN` field has been added to the `publication` type.  Why on earth this isn't standard,
      I have no idea!

* This utility **also supports an XML representation** of the JSON Resume data schema, if you prefer
  the XML format or tooling.  The utility can convert a resume data file back and forth between
  both formats.

### The Bad?

* HackMyResume comes with a bundle of resume templates out-of-the box, and there may be more
  created by the community.  I'm a one-man show, and **the number of templates is much more limited**.
  Templates from HackMyResume can't be used as-is here... because they are based on the
  [Handlebars](http://handlebarsjs.com/) template processor, while this utility uses
  [Go templates](https://golang.org/pkg/text/template/).  That being said, the two template formats
  are very similar, and it shouldn't take *extraordinary* effort to convert them.

* HackMyResume allows you to generate resumes in multiple output formats (e.g. Microsoft Word,
  PDF, HTML, Markdown, etc).  My utility **only publishes to the Word format**, period.  To me it
  doesn't make sense to force all the effort of creating multiple versions of each template to support
  all possible outputs.  In reality, 99% of the time people want your resume in Word format... and
  Word (or LibreOffice) can always export to PDF or HTML anyway.

* HackMyResume is supposed to be working toward support for HTML or Markdown formatting within
  your resume *data* file.  At first, I was too.  That's why I wrote support for both JSON and
  XML data files, since XML is more friendly toward line-breaks within field values.
  However, I found this effort to be really difficult, and so **I've backed off from HTML/Markdown
  field support for now**.  The Microsoft Word file format is a mess... inserting arbitrary *text* is
  easy, but inserting arbitrary *formatting* is a nightmare.  If HackMyResume or someone else comes up
  with an elegant approach, then I might revisit this.  However, the optional schema extensions that I
  added have pretty much eliminated my own personal need for HTML/Markdown formatting support.

* The developer behind HackMyResume has been a lot more active, and seems more committed to
  building and supporting a community.  In contrast, this processor is a personal project that
  I wrote mainly for my own needs.  I've made it public, just to see if others do anything with
  it, but **I don't have the bandwidth to update this regularly or provide much support**.  So if
  community is a big deal to you, then HackMyResume is awesome and I definitely recommend it.
  However, if you want to use some of my schema extensions, or if you just love Go or hate Node,
  then welcome aboard!  At the end of the day, you're just talking about a resume file.  If your
  preferred generator tool disappears, or turns evil, or whatever... you can always just
  cut-and-paste your content into the next thing.

Using the Processor
-------------------
```
Usage:

   resume COMMAND <args>

... where "COMMAND" is one of the following:

	init    - Create a new empty resume data file
	convert - Convert a JSON-formatted resume data file into XML
	          format, or vice-versa
	export  - Process and resume data file with a given template,
	          to publish a Microsoft Word resume file

Full details for each command:

resume init <filename>
resume init resume.xml

	Will generate an empty resume data file with the specified
	filename, which must have either a '.json' or '.xml' file
	extension.

	If no filename is specified, then a data file will be created
	with filename 'resume.json'.

resume convert <input filename> <output filename>
resume convert resume.xml resume.json

	The resume data file specified by the first parameter will
	be converted to the filename specified by the second parameter.

	If the second file already exists, then any contents will
	be overwritten.  Both filenames must have either a '.json' or
	'.xml' file extension.

resume export <data filename> <output filename> <template filename>
resume export resume.json resume.doc templates/plain.xml

	The resume data file specified by the first parameter will
	published as a Microsoft Word file with the name specified by
	the second parameter.  The template file specified by the third
	parameter will be used to generate the output.

	The data filename must have either a '.json' or '.xml' file
	extension.  The output will be a Microsoft Word 2003 XML file,
	and its name must have either a '.doc' or '.xml' file
	extension.  The template file must likewise be a Word 2003
	XML file (with Go template tags), and its name too must have
	either a '.doc' or '.xml' extension.

	If the specified template is not found in the current working
	directory, then the application will look under a "templates"
	subdirectory in the current working directory.  If no template
	is specified, the the application will use the "plain.xml"
	template.
```

Creating new templates
----------------------
Just like with HackMyResume, templates here are Microsoft Word 2003 XML files, with text values
inserted by a template engine.

Although a bit arcane, the Word 2003 XML format is actually quite thoroughly documented:

* https://en.wikipedia.org/wiki/Microsoft_Office_XML_formats
* https://www.microsoft.com/en-us/download/details.aspx?id=101

However, you can probably figure out most of what you need by simply reviewing the existing
template files under the `templates` subdirectory.

The big difference between this and HackMyResume is that this uses
[Go templates](https://golang.org/pkg/text/template/), rather than the
[Handlebars](http://handlebarsjs.com/) template engine.  There are some differences, but both
engines essentially involve directives and values wrapped by double-curly-braces.

Handlebars:
```
{{#each r.employment.history}}
{{ position }}
{{/each}}
```

Go templates:
```
{{ range .Work }}
{{ .Position }}
{{ end }}
```

If you have no experience (or at least interest) in Go, then creating new templates for this
processor may be a daunting task.  Otherwise, the process is fairly straightforward:

1. Use Microsoft Word (or LibreOffice) to create a resume with whatever style you like.

2. Specify the "Word 2003 XML Document" file type when saving the file.  It's very important
   to be precise about this, as there are multiple different Word file options with "2003"
   or "XML" in their name.

3. Open the file in a text editor.  For all dynamic fields, replace your sample text with Go
   template tags that reference the appropriate JSON Resume schema fields.  Resources to help
   you are:

   * Existing template files, under the `templates` subdirectory.

   * The [documentation](https://golang.org/pkg/text/template/) for Go templates.

   * The data schema structures in the `data/data.go` file.

4. Save your new template file under the `templates` subdirectory.

A Note About File Formats (do you need MS Word?)
------------------------------------------------
This processor, like HackMyResume, produces resume output files in Microsoft Word 2003 XML
format.  Every modern version of Microsoft Word handles these files great, so if you'll always
be working with Word then you can stop reading here!

However, what if don't have access to Microsoft Word on your machine (e.g. you're on Linux,
you have strong free software beliefs, you simply don't want to buy it, whatever)?  Unfortunately,
the 2003 XML format variation is not well supported by LibreOffice, or online options such as
Google Docs or Word Online.

Fortunately, the most recent default Word format IS well-supported by LibreOffice and Word
Online (I actually think that LibreOffice does a better job of exporting Word docs to PDF than
Word itself does).  So if you want, you can convert your resume output to the latest Word format
through one of these options:

1. If you have access to a computer with Microsoft Word installed, then obviously you can open
   your resume on that machine and "Save As..." to a copy with the default latest format.

2. Otherwise, I've had nothing but great results with
   [CloudConvert](https://cloudconvert.com/doc-to-docx), a free online service that (among other
   things) supports conversion from DOC to DOCX.
