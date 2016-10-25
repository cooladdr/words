var Word = React.createClass({
    render:function(){
        return (
			<div>
				<div className='word'>{this.props.word.Spelling}</div>
				<div className='Meaning'>{this.props.word.Definitions}</div>
			</div>
            
        );
    }
});
var WordsShow = React.createClass({

    render:function(){
        var words = this.props.words.map(function(word){
            return (
                <Word word={word} url=''/>
            );
        });

        if(this.props.show_status){
            return (
                <div className='words-show-block'>{words}</div>
            );
        }else{
            return (
                <div className='words-show-block'></div>
            );
        }
    }
});

var Book = React.createClass({
    getInitialState: function() {
        return {
            show_status:true,
            words:[]
        };
    },
    handleClick:function(){
        if(this.state.words.length > 0){
            this.setState({show_status: !this.state.show_status});
        }else{
            $.ajax({
                url: this.props.url,
                dataType: 'json',
                cache: false,
                data:{id:this.props.book.id},
                success: function(data) {
                    this.setState({words: data});
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }
    },
	redirect:function(){
		//location.href='/book/'+this.props.book.Id;
		window.open('/book/'+this.props.book.Id);
	},
    render:function(){
        return (
            <div className='book-row'>
                <h4 className='book' onClick={this.redirect}>{this.props.book.Name}</h4>
				<div className='note'>{this.props.book.Count}词</div>
                <WordsShow words={this.state.words} show_status={this.state.show_status}/>
				
				<hr  className='hLine'/>
            </div>
        );
    }
});

var BooksShow = React.createClass({
    render:function(){
        var books = this.props.books.map(function(book){
            return (
                <Book book={book} url={'/api/cat/book/'+book.Id}/>
            );
        });

        if(this.props.show_status){
            return (
                <div className='books-show-block'>{books}</div>
            );
        }else{
            return (
                <div className='books-show-block'></div>
            );
        }

    }
});
var Category = React.createClass({
    getInitialState: function() {
        return {
            show_status:true,
            books:[]
        };
    },

    handleClick:function(){
        if(this.state.books.length > 0){
            this.setState({show_status: !this.state.show_status});
        }else{
            $.ajax({
                url: this.props.url,
                dataType: 'json',
                cache: false,
                data:{id:this.props.catData.id},
                success: function(data) {
                    this.setState({books: data});
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }
    },
    render : function(){
        return (
            <div className = 'category-row' >
                <h3 className='category' onClick={this.handleClick}>{this.props.catData.Name}<span className='note'>{this.props.catData.Count}本</span></h3>
                <BooksShow books={this.state.books} show_status={this.state.show_status}/>
                <hr className='hLine' />
            </div>
        );
    }
});

var CategoryShow = React.createClass({
    getInitialState: function() {
        return {
            categories:[]
        };
    },

    componentDidMount:function(){
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            cache: false,
            success: function(data) {
                this.setState({categories: data});
            }.bind(this),
            error: function(xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },

    render : function(){
        var categories = this.state.categories.map(function(category){
            return (
                <Category catData={category} url={'/api/cat/'+category.Id}/>
            );
        });
        return (
            <div className='category-show-block'>
            {categories}
            </div>
        );
    }
});


ReactDOM.render(
    <CategoryShow url = '/api/cat'/>,
    document.getElementById('container')
);