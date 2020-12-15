---
title: "Emacs - Adding clipboad.js to HTML exports"
date: 2019-09-18
author: Riz
categories: ["digital"]
tags: ["emacs", "linux"]
keywords: ["emacs", "clipboard", "codeblock", "clipboardjs", "export", "html"]
---


[Clipboard.js](https://clipboardjs.com/) is a javascript solution to add copy buttons to `<pre>` or `<code>` regions in your html file. F

Since many of the Emacs community use org-publish for blogging, it seems like a good idea to note how to include clipboard.js snippets in the org mode html exports.

-   Add the following to your `init.el`
```lisp
(defun org-html-src-block (src-block contents info)
      "Transcode a SRC-BLOCK element from Org to HTML.
    CONTENTS holds the contents of the item.  INFO is a plist holding
    contextual information."
      (if (org-export-read-attribute :attr_html src-block :textarea)
          (org-html--textarea-block src-block)
        (let* ((lang (org-element-property :language src-block))
    	  (caption (org-export-get-caption src-block))
    	  (code (org-html-format-code src-block info))
    	  (label (let ((lbl (and (org-element-property :name src-block)
    				 (org-export-get-reference src-block info))))
    		   (if lbl (format " id=\"%s\"" lbl)
    		     (concat "btn_" (s-replace "-" "" (org-id-new))))))
    	  (btnprops " class='btn-clipboard'  data-original-title='Copy to clipboard'")
    	  (button (concat "<div class=\"bd-clipboard\"> <button" btnprops " data-clipboard-target=\"#" label "\"> Copy </button> </div>")))
          (if (not lang) (format "<pre class=\"example\"%s>\n%s</pre>" label code)
    	(format
    	 "<div class=\"org-src-container\">\n%s%s\n</div>"
    	 (if (not caption) ""
    	   (format "<label class=\"org-src-name\">%s</label>"
    		   (org-export-data caption info)))
    	 (format "%s\n<pre class=\"src src-%s\" id=\"%s\">%s</pre>" button lang label code))))))
    
    
    (defconst org-clipboardjs-script
    "<!-- 2. Include library -->
        <script src=\"https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.0/clipboard.min.js\"></script>
    
        <!-- 3. Instantiate clipboard by passing a string selector -->
        <script>
        var clipboard = new ClipboardJS('.btn-clipboard');
        clipboard.on('success', function(e) {
    	console.log(e);
        });
        clipboard.on('error', function(e) {
    	console.log(e);
        });
        </script>"
    "Clipoard JS script to be included"
    )
    
    
    (defun org-html-template (contents info)
      "Return complete document string after HTML conversion.
    CONTENTS is the transcoded contents string.  INFO is a plist
    holding export options."
      (concat
       (when (and (not (org-html-html5-p info)) (org-html-xhtml-p info))
         (let* ((xml-declaration (plist-get info :html-xml-declaration))
    	    (decl (or (and (stringp xml-declaration) xml-declaration)
    		      (cdr (assoc (plist-get info :html-extension)
    				  xml-declaration))
    		      (cdr (assoc "html" xml-declaration))
    		      "")))
           (when (not (or (not decl) (string= "" decl)))
    	 (format "%s\n"
    		 (format decl
    			 (or (and org-html-coding-system
    				  (fboundp 'coding-system-get)
    				  (coding-system-get org-html-coding-system 'mime-charset))
    			     "iso-8859-1"))))))
       (org-html-doctype info)
       "\n"
       (concat "<html"
    	   (when (org-html-xhtml-p info)
    	     (format
    	      " xmlns=\"http://www.w3.org/1999/xhtml\" lang=\"%s\" xml:lang=\"%s\""
    	      (plist-get info :language) (plist-get info :language)))
    	   ">\n")
       "<head>\n"
       (org-html--build-meta-info info)
       (org-html--build-head info)
       (org-html--build-mathjax-config info)
       "</head>\n"
       "<body>\n"
       (let ((link-up (org-trim (plist-get info :html-link-up)))
    	 (link-home (org-trim (plist-get info :html-link-home))))
         (unless (and (string= link-up "") (string= link-home ""))
           (format (plist-get info :html-home/up-format)
    	       (or link-up link-home)
    	       (or link-home link-up))))
       ;; Preamble.
       (org-html--build-pre/postamble 'preamble info)
       ;; Document contents.
       (let ((div (assq 'content (plist-get info :html-divs))))
         (format "<%s id=\"%s\">\n" (nth 1 div) (nth 2 div)))
       ;; Document title.
       (when (plist-get info :with-title)
         (let ((title (plist-get info :title))
    	   (subtitle (plist-get info :subtitle)))
           (when title
    	 (format
    	  (if (plist-get info :html-html5-fancy)
    	      "<header>\n<h1 class=\"title\">%s</h1>\n%s</header>"
    	    "<h1 class=\"title\">%s%s</h1>\n")
    	  (org-export-data title info)
    	  (if subtitle
    	      (format
    	       (if (plist-get info :html-html5-fancy)
    		   "<p class=\"subtitle\">%s</p>\n"
    		 "\n<br>\n<span class=\"subtitle\">%s</span>\n")
    	       (org-export-data subtitle info))
    	    "")))))
       contents
       (format "</%s>\n" (nth 1 (assq 'content (plist-get info :html-divs))))
       ;; Postamble.
       (org-html--build-pre/postamble 'postamble info)
    
       (org-element-normalize-string org-clipboardjs-script)
       ;; Closing document.
       "</body>\n</html>"))
  ```
-   Reload your configuration with `M-x eval-buffer` on init.el
-   Now export any org file with `M-x org-html-export-to-html` or with shortcut keys <kbd>C-c C-e h h</kbd>
