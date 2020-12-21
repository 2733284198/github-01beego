$(function(){	
	app.init()	
})

var app={
	init:function(){
		this.slideToggle();
		this.resizeIframe()
	},	
	slideToggle:function(){
		$('.aside h4').click(function(){
		
			$(this).siblings('ul').slideToggle();
		})
	},
	resizeIframe:function(){		
		$("#rightMain").height($(window).height()-80)
	}
}
