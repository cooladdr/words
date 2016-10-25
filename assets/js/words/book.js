var Word = React.createClass({
    render:function(){
        return (
			<div>
				<div className='word'>{this.props.word.Spelling}</div>
				<div className='Meaning'>{this.props.word.Definitions}</div>
				<hr  className='hLine'/>
			</div>
            
        );
    }
});


var WordsShow = React.createClass({

	getInitialState: function() {
		var bookId = document.getElementById('bookId').className;
        return {
			bookId:bookId,
			page:1,
            words:[],
			hasMore:true
        };
    },
	
	getWords:function(){
		$.ajax({
            url: this.props.url+this.state.bookId,
            dataType: 'json',
            cache: false,
			data:{page:this.state.page, size:20},
            success: function(data) {
				if(data.length == 20){
					this.setState({words: this.state.words.concat(data), hasMore:true, page:this.state.page+1});
				}else{
					this.setState({words: this.state.words.concat(data), hasMore:false});
				}
                
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
	},
	
	getMore:function(){
		if(this.state.hasMore){
			this.getWords();
		}
	},

    componentDidMount:function(){
		this.getWords();
    },
	
	scrollState: function(scroll) {
		alert(scroll);
	    var visibleStart = Math.floor(scroll / this.state.recordHeight);
	    var visibleEnd = Math.min(visibleStart + this.state.recordsPerBody, this.state.total - 1);
	
	    var displayStart = Math.max(0, Math.floor(scroll / this.state.recordHeight) - this.state.recordsPerBody * 1.5);
	    var displayEnd = Math.min(displayStart + 4 * this.state.recordsPerBody, this.state.total - 1);
	
	    this.setState({
	        visibleStart: visibleStart,
	        visibleEnd: visibleEnd,
	        displayStart: displayStart,
	        displayEnd: displayEnd,
	        scroll: scroll
	    });
	},

    render:function(){
        var words = this.state.words.map(function(word){
            return (
                <Word word={word} url=''/>
            );
        });
		
		if(this.state.hasMore){
			return (
				<div>
	            <div className='words-show-block'>{words}</div>
				<div className='get-more' onClick={this.getMore}>点击获取更多</div>
				<hr className='hLine'/>
				</div>
	        );
		}else{
			return (
				<div>
	            <div className='words-show-block'>{words}</div>
				</div>
	        );
		}
		
        
        
    }
});




ReactDOM.render(
    <WordsShow url = '/api/cat/book/'/>,
    document.getElementById('container')
);