{{define "navbar"}}
<a class="navbar-brand" href="/">庫存管理系統</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .Import}}class="active"{{end}}><a href="/import">進貨明細</a></li>
        <li {{if .Export}}class="active"{{end}}><a href="/export">出貨明細</a></li>
        <li {{if .Increase}}class="active"{{end}}><a href="/increases">新增管理</a></li>
        <li {{if .WareHouse}}class="active"{{end}}><a href="/warehouse">庫存管理</a></li>
    </ul>
</div>
<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
		<li><a href="/login?exit=true">退出</a></li>
		{{else}}
		<li><a href="/login">會員管理員</a></li>
		{{end}}
	
</div>
{{end}}