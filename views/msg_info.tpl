<div class='alert alert-[[.Tipo]] alert-dismissible' role='alert'>
    <button type='button' class='close' data-dismiss='alert' aria-label='Close'>
	    <span aria-hidden='true'>&times;</span>
	</button>
	<ul>
	    [[with .Erros]]
		[[range .]]<li><strong>[[.Key]]</strong> [[.Message]]</li>[[end]]
		[[end]]
		[[if .Single]]
		<li><strong>[[.Key]]</strong> [[.Message]]</li>
		[[end]]
	</ul>
</div>