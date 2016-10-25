
var WordRow = React.createClass({
    render: function() {
		if(this.props.word_info){
			return (
	            <div className='word-row'>
	                <span>
	                    <span className='spelling'>{this.props.word}</span>
	                    {this.props.word_info.Meiyin ? <span className='pronunciation'><span>美 : </span>[  {this.props.word_info.Meiyin}  ]</span> : <span></span>}
	                    {this.props.word_info.Yingyin ? <span className='pronunciation'><span>英 : </span>[  {this.props.word_info.Yingyin}  ]</span> : <span></span>}
	                </span>
	            </div>
	        );
		}else{
			return (
	            <div className='word-row'></div>
	        );
		}
        
    }
});


var WordInfo = React.createClass({
    render: function() {
        return (
            <div className='word-info'>
                <WordRow word={this.props.word} word_info={this.props.word_info}/>
            </div>
        );
    }
});

var SearchBar = React.createClass({
    handleChange: function() {
        this.props.onUserInput(
            this.refs.filterTextInput.value
        );
    },
    handleBlur:function(){
        //this.props.onUserBlur();
    },
    handleKeyPress:function(event){
        if(event.keyCode == 13) {
            //this.props.onUserBlur();
        }
    },
    render: function() {
        return (
            <div className='search-bar'>
                <input type="text"
                placeholder="输入单词"
                value={this.props.word}
                ref="filterTextInput"
                onChange={this.handleChange}
                onBlur={this.handleBlur}
                onKeyDown={this.handleKeyPress}
                />
            </div>
        );
    }
});

var Examples = React.createClass({
    render:function(){
        if(this.props.show_status){
			var expls = "";
			if(this.props.data){
				expls = this.props.data.map(function(ex){
		            return (
						<div><p>{ex.EnSentence}</p>
						<p>{ex.ChSentence}</p></div>
					);
		        });
			}
            return (<div className='example'>{expls}</div>);
        }else{
            return (<div className='example'></div>);
        }

    }
});

var Definition = React.createClass({
    getInitialState: function() {
        return {
            show_status:true,
            example_done:false,
            examples: []
        };
    },
    handleClick:function(){
        if(this.state.example_done){
            this.setState({show_status: !this.state.show_status});
        }else{
            $.ajax({
                url: '/api/word/expl/'+this.props.def.Id,
                dataType: 'json',
                cache: false,
                data:{},
                success: function(data) {
					if(data){
						this.setState({examples: data, example_done:true});
					}
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }
    },
    render:function(){
        return <div className='def-row'>
            <span className='part'>{this.props.def.Part != 'etymon' ? this.props.def.Part : '.'}&nbsp;&nbsp;&nbsp;</span>
            <span>{this.props.def.ChDef}</span>
            <span className='show-example' onClick={this.handleClick}>例子&gt;&gt;</span>
            <Examples data={this.state.examples} show_status={this.state.show_status}/>
        </div>
    }
});

var DefinitionShow = React.createClass({
    render:function(){
        var defs = "";
		if(this.props.defs){
			defs = this.props.defs.map(function(def){
	            return <Definition def={def} />
	        });
		}
		
        return (
            <div className='DefinitionShow'>{defs}</div>
        );
    }
});

var AWord = React.createClass({
    handleWordClick:function(){
        this.props.onLiClick(this.props.relData.Spelling);
    },
    render : function(){
        return (
            <span className='word' onClick={this.handleWordClick}>{this.props.relData.Spelling}</span>
        );
    }
});

var RelationshipShow = React.createClass({
    relationMap : function(needle){
        var map = {
            'adjcomp':'形容词比较级',
            'adjsuper':'形容词最高级',
            'advcomp':'副词比较级',
            'advsuper':'副词最高级',
            'pasttense':'过去式',
            'pastparticiple':'过去分词',
            'presentparticiple':'现在分词',
            'singular':'第三人称单数',
            'antonym':'反义词',
            'synonym':'同(近)义词',
            'derivative':'派生词',
            'paronym':'同源词',
            'similar':'形近词',
            'phrase':'相关短语',
            'plural':'名词复数',
            'compound':'组合词'
        }
        return map[needle];
    },
    handleClick:function(word){
        this.props.onUserInput(word);
    },
    render : function(){
        var rows=[];
        var last_rel=null;
        var children=null;
        this.props.rels.forEach(function(rel){
            if(rel.Relation != last_rel){
                if(children){
                    rows.push(
                        <div className='relation-row'>
                            <span className='relation-title'>{this.relationMap(last_rel)}: </span>
                            {children}
							<div className='floatClean'></div>
                        </div>
                    );
                }
                last_rel = rel.Relation;
                children=[];
            }
            children.push(<AWord onLiClick={this.handleClick} relData={rel}/>)
        }.bind(this));

        if(rows.length){
            return (
                <div className='relationship'>
				<hr className='hLine' />
                    {rows}
                </div>
            );
        }else{
            return (
                <div className='relationship'></div>
            );
        }

    }
});

var MainBox = React.createClass({
    getInitialState: function() {
        return {
            word: '',
            info:[],
            defs:[],
            rels:[]
        };
    },

    handleUserInput: function(word) {
		this.setState({word: word, info:[],defs:[],rels:[]});
        if(word){
            $.ajax({
                url: this.props.url+word,
                dataType: 'json',
                cache: false,
                data:{word:word},
                success: function(data) {
					if(data){
						this.setState({info: data});
						this.getDefinitions();
						this.getRelationship();
					}
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }else{
            this.setState({info: []});
        }
    },

    getDefinitions:function(){
		//alert('getDefinitions');
        if(this.state.word){
            $.ajax({
                url: '/api/word/def/'+this.state.info.Id,
                dataType: 'json',
                cache: false,
                data:{},
                success: function(data) {
					if(data){
						this.setState({defs: data});
					}
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }else{
            this.setState({defs: []});
        }
    },

    getRelationship:function(){
        if(this.state.word){
            $.ajax({
                url: '/api/word/'+this.state.word+'/relation',
                dataType: 'json',
                cache: false,
                data:{word:this.state.word},
                success: function(data) {
					if(data){
						this.setState({rels: data});
					}
                }.bind(this),
                error: function(xhr, status, err) {
                    console.error(this.props.url, status, err.toString());
                }.bind(this)
            });
        }else{
            this.setState({rels: []});
        }
    },
    render: function() {
        return (
            <div classname='main-box' >
                <SearchBar word={this.state.word} onUserInput={this.handleUserInput}/>
                <WordInfo word={this.state.word} word_info={this.state.info} />
                <DefinitionShow defs={this.state.defs} />
                <RelationshipShow word={this.state.word} rels={this.state.rels}  onUserInput={this.handleUserInput}/>
            </div>
        );
    }
});

ReactDOM.render(
    <MainBox url = '/api/word/'/>,
    document.getElementById('container')
);