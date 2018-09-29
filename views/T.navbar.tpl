{{define "navbar"}}
<nav class="navbar navbar-default">

    <ul class="nav navbar-nav">

      <li {{if .IsStandard}}class="active"{{end}}>
        <a href="/standard">规范</a>
      </li>
      <li {{if .IsLegis}}class="active"{{end}}>
        <a href="/legislation">对标</a>
      </li>
    </ul>

</nav>
{{end}}