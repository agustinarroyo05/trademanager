function TransactionList(props) {
	var txList = null;

	function getTxs(data){
		if (data === null)
			return;
		const listItems = data.map((tx) =>
			<Transaction amount={tx.amount} type={tx.type} id={tx.id} efectiveDate={tx.efectiveDate}/>
		);
		txList = <ul className='list-group,table, ulwrap'>{listItems}</ul>;
	}
	$.ajax({url : 'transactiontrade/transactions',
	  success: getTxs,
	  async: false
	});
	
	return txList;
}

class Transaction extends React.Component{

  constructor(props) {
    super(props);
    this.state = {isToggleOn: false, amount : props.amount, type : props.type, id : props.id, efectiveDate : props.efectiveDate};

    this.showDetailsBtn = this.showDetailsBtn.bind(this);
  }

  showDetailsBtn(){
   this.setState(state => ({
      isToggleOn: !state.isToggleOn
    }));
  }
  render(){
	  if (this.state.isToggleOn){
		  return(<li className='list-group-item'> 	
		  	<button onClick={this.showDetailsBtn}>{this.state.isToggleOn ? '-' : '+'}</button>
		  	<span className='ppal'>Type: {this.state.type === 1 ? 'DEBIT' : 'CREDIT'}</span> 
		  	<span className={this.state.type === 1 ? 'negcred' : ''} > Amount: {this.state.amount}</span>
		  	<div className='details'>
				<span> Id: {this.state.id} </span>
				<span> Time : {this.state.efectiveDate} </span>
			</div>
		 </li>);
		}
		else
			return(<li className='list-group-item'> 	
				<button onClick={this.showDetailsBtn}>{this.state.isToggleOn ? '-' : '+'}</button>
			  	<span className='ppal'>Id: {this.state.id}</span> 
			  	<span className={this.state.type === 1 ? 'negcred' : ''} > Amount: {this.state.amount}</span>
		 </li>);

	}
}

ReactDOM.render(
  <TransactionList />,
  document.getElementById('root')
);