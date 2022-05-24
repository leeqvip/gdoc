package views

func Index() string {

	return `
{{ define "gloc" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>{{ .title }}</title>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style type="text/css">
        {{ .css }}
    </style>
</head>

<body>
<div id="container">
    <header id="header" class="wrapper">
        <div id="header-inner" class="inner">
            <h1 id="logo-wrap">
                <span>Documentation</span>
            </h1>
            <nav id="main-nav">
            </nav>
            <div id="lang-select-wrap">
				<label id="lang-select-label"><i class="fa fa-globe"></i><span>{{ .current_version_title }}</span></label>
                <select id="lang-select" data-canonical="">
					{{ range $key, $version := .versions }}
                    <option value="{{$key}}"
					{{ if eq $.current_version  $key }} 
						selected 
					{{ end }}
					>{{$version}}</option>
                    {{ end }}
                </select>
                
            </div>
            <a id="mobile-nav-toggle">
                <span class="mobile-nav-toggle-bar"></span>
                <span class="mobile-nav-toggle-bar"></span>
                <span class="mobile-nav-toggle-bar"></span>
            </a>
        </div>
    </header>
    <div id="content-wrap">
        <div id="content" class="wrapper">
            <div id="content-inner">
                <article class="article-container" itemscope itemtype="http://schema.org/Article">
                    <div class="article-inner">
                        <div class="article">
                            <div class="inner">
                                <!-- <header class="article-header">
                                    <h1 class="article-title" itemprop="name">--</h1>
                                </header> -->
                                <div class="article-content" itemprop="articleBody">
                                    {{ .content }}
                                </div>
                                <footer class="article-footer">
                                    <!-- <time class="article-footer-updated" datetime="2019-03-26T07:54:33.316Z" itemprop="dateModified">上次更新：2019-03-26</time> -->
                                </footer>
                            </div>
                        </div>
    					<aside id="article-toc" role="navigation">
							<div id="article-toc-inner">                                
								<strong class="sidebar-title">目录</strong>							
								<a href="#" id="article-toc-top">回到顶部</a>
							</div>
                        </aside>
                    </div>
                </article>
                <aside id="sidebar" role="navigation">
                    <div class="inner"></div>
                </aside>
            </div>
        </div>
    </div>
    <footer id="footer" class="wrapper">
        <div class="inner">
            <div id="footer-copyright">
                &copy; 2019 Ldocs
            </div>
            <div id="footer-links">
            </div>
        </div>
    </footer>
</div>

<div id="mobile-nav-dimmer"></div>
<nav id="mobile-nav">
    <div id="mobile-nav-inner">
        <!-- <ul id="mobile-nav-list">

          <li class="mobile-nav-item">
            <a href="https://github.com/" class="mobile-nav-link" rel="external" target="_blank">GitHub</a>
          </li>
        </ul> -->
    </div>
    <div id="mobile-lang-select-wrap">
        <span id="mobile-lang-select-label"><i class="fa fa-globe"></i><span>{{ .current_version_title }}</span></span>
        <select id="mobile-lang-select" data-canonical="">
            {{ range $key, $version := .versions }}
                    <option value="{{$key}}"
					{{ if eq $.current_version  $key }} 
						selected 
					{{ end }}
					>{{$version}}</option>
                    {{ end }}
        </select>
    </div>
</nav>

<script src="https://cdn.bootcdn.net/ajax/libs/esprima/2.7.3/esprima.min.js"></script>
<script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
<script src="https://cdn.bootcdn.net/ajax/libs/js-yaml/4.1.0/js-yaml.min.js"></script>
<script type="text/javascript">
    var sidebar = jsyaml.load(` + "`{{ .sidebar }}`" + `);
    var sidebarHtml = ''
    var mobileSidebarHtml = ''
    var basePath = '{{ .basePath }}'
    var contentFileName = '{{ .contentFileName }}'

    for(var children in sidebar){
        sidebarHtml += '<strong class="sidebar-title">' + children + '</strong>';
        mobileSidebarHtml += '<strong class="mobile-nav-title">' + children + '</strong>';
        for(var child in sidebar[children]){
            current = contentFileName == sidebar[children][child] ? 'current' : '';
            sidebarHtml += '<a href="'+basePath + sidebar[children][child] + '" class="sidebar-link toc-link '+current+'">'+child+'</a>';
            mobileSidebarHtml += '<a href="'+basePath + sidebar[children][child] + '" class="mobile-nav-link toc-link '+current+'">'+child+'</a>';
        }
    }
    $("#sidebar .inner").html(sidebarHtml);
    $("#mobile-nav-inner").html(mobileSidebarHtml);

    function toc(str, options = {}) {
        const headingsMaxDepth = options.hasOwnProperty('max_depth') ? options.max_depth : 2;
        const headingsSelector = ['h2', 'h3', 'h4', 'h5', 'h6'].slice(0, headingsMaxDepth).join(',');

        const headings = $('<div>' + str + '<div>').find(headingsSelector);
        if (!headings.length) return '';

        const className = options.class || 'toc';
        const listNumber = options.hasOwnProperty('list_number') ? options.list_number : false;
        let result = '<ol class="' + className + '">';
        const lastNumber = [0, 0, 0, 0, 0, 0];
        let firstLevel = 0;
        let lastLevel = 0;

        function getId(ele) {
            const id = ele.attr('id');
            const $parent = ele.parent();
            return id ||
                ($parent.length < 1 ? null :
                    getId($parent));
        }

        headings.each(function(index, el) {
            const level = +$(this).prop("localName")[1];
            const id = getId($(this));
            const text = $(this).text();

            lastNumber[level - 1]++;

            for (let i = level; i <= 5; i++) {
                lastNumber[i] = 0;
            }

            if (firstLevel) {
                for (let i = level; i < lastLevel; i++) {
                    result += '</li></ol>';
                }

                if (level > lastLevel) {
                    result += '<ol class="' + className + '-child">';
                } else {
                    result += '</li>';
                }
            } else {
                firstLevel = level;
            }

            result += '<li class="' + className + '-item ' + className + '-level-' + level+ '">';
            result += '<a class="' + className + '-link" href="#' + id+ '">';

            if (listNumber) {
                result += '<span class="' + className + '-number">';

                for (let i = firstLevel - 1; i < level; i++) {
                    result += lastNumber[i] + '.';
                }

                result += '</span> ';
            }

            result += '<span class="' + className + '-text">' + text+ '</span></a>';

            lastLevel = level;
        });

        for (let i = firstLevel - 1; i < lastLevel; i++) {
            result += '</li></ol>';
        }

        return result;
    }
    articleToc = toc($(".article-content").html());
    $("#article-toc-inner .sidebar-title").after(articleToc);

	(function() {
			'use strict';


			function changeLang() {
				var lang = this.value;
				var canonical = this.dataset.canonical;
				var path = "{{ .prefix_uri }}" ;
				if(lang != '{{ .default_version_name }}') path += lang + '/';
				location.href = path + canonical;
			}

			document.getElementById('lang-select').addEventListener('change', changeLang);
			document.getElementById('mobile-lang-select').addEventListener('change', changeLang);
		}());
		(function() {
			'use strict';

			var body = document.getElementsByTagName('body')[0];
			var navToggle = document.getElementById('mobile-nav-toggle');
			var dimmer = document.getElementById('mobile-nav-dimmer');
			var CLASS_NAME = 'mobile-nav-on';
			if (!navToggle) return;

			navToggle.addEventListener('click', function(e) {
				e.preventDefault();
				e.stopPropagation();
				body.classList.toggle(CLASS_NAME);
			});

			dimmer.addEventListener('click', function(e) {
				if (!body.classList.contains(CLASS_NAME)) return;

				e.preventDefault();
				body.classList.remove(CLASS_NAME);
			});
		}());

</script>
</body>

</html>
{{ end }}
`
}
