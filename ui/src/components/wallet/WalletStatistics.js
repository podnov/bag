import React, { Component } from 'react'
import {connect} from 'react-redux'
import equal from 'fast-deep-equal'
import Loader from 'react-loader-spinner'

class WalletStatistics extends Component {
	constructor(props) {
		super(props);

		this.state = {
			isLoading: this.props.isLoading,
			statistics: this.props.statistics
		}
	}

	componentDidUpdate(prevProps) {
		console.log('comp did update');
		if(!equal(this.props, prevProps)) {
			console.log('comp did update change');
			this.setState({
				isLoading: this.props.isLoading,
				statistics: this.props.statistics
			});
		}
	}

	render() {
		console.log('render');
		const isLoading = this.state.isLoading;
		const statistics = this.state.statistics;

		let content;

		if (isLoading) {
			console.log('render loading');
			content = <Loader />;
		} else if (statistics) {
			console.log('render statistics');
			content = <div className="walletStatistics">
				<h2>Summary</h2>
				<ul>
					<li><label>Address:</label> {statistics.accountAddress}</li>
					<li><label>First Transaction:</label> {statistics.firstTransactionAt}</li>
					<li><label>Transaction Count:</label> {statistics.transactionCount}</li>
					<li><label>Value:</label> {statistics.value}</li>
					<li><label>Earned Value:</label> {statistics.earnedValue}</li>
					<li><label>Earned Value Per Day:</label> {statistics.earnedValuePerDay}</li>
					<li><label>Earned Value Per Week:</label> {statistics.earnedValuePerWeek}</li>
				</ul>
				<h2>Tokens</h2>
				{statistics.tokens.map(function(token, index) {
					return (
						<div>
							<h3>Token: {token.tokenName} ({token.tokenAddress})</h3>
							<ul>
								<li><label>First Transaction:</label> {token.firstTransactionAt}</li>
								<li><label>Transaction Count:</label> {token.transactionCount}</li>
								<li><label>Price:</label> {token.tokenPrice}</li>
								<li><label>Price Updated At:</label> {token.tokenPriceUpdatedAt}</li>
								<li><label>Count:</label> {token.tokenCount}</li>
								<li><label>Value:</label> {token.value}</li>
								<li><label>Earned Count:</label> {token.earnedTokenCount}</li>
								<li><label>Earned Count Per Day:</label> {token.earnedTokenCountPerDay}</li>
								<li><label>Earned Count Per Week:</label> {token.earnedTokenCountPerWeek}</li>
								<li><label>Earned Value:</label> {token.earnedValue}</li>
								<li><label>Earned Value Per Day:</label> {token.earnedTokenValuePerDay}</li>
								<li><label>Earned Value Per Week:</label> {token.earnedTokenValuePerWeek}</li>
							</ul>
						</div>
					)
				})}
			</div>;
		} else {
			console.log('render else');
			content = <div>Welcome, please enter your wallet address</div>;
		}

		return (
			<div>
				{content}
			</div>
		)
	}
}

const mapStateToProps = (state) => ({
	isLoading: state.wallet.isLoading,
	statistics: state.wallet.statistics
})

export default connect(mapStateToProps)(WalletStatistics)

