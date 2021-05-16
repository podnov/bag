import React, { Component } from 'react'
import {connect} from 'react-redux'
import equal from 'fast-deep-equal'
import Loader from 'react-loader-spinner'

import './WalletStatistics.css'

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
			content = <div>
				<h2>Summary</h2>
				<ul>
					<li key="summaryAddress"><label>Address:</label> {statistics.accountAddress}</li>
					<li key="summaryFirstTransactionAt"><label>First Transaction:</label> {statistics.firstTransactionAt}</li>
					<li key="summaryTransactionCount"><label>Transaction Count:</label> {statistics.transactionCount}</li>
					<li key="summaryValue"><label>Value:</label> {statistics.value}</li>
					<li key="summaryEarnedValue"><label>Earned Value:</label> {statistics.earnedValue}</li>
					<li key="summaryEarnedValuePerDay"><label>Earned Value Per Day:</label> {statistics.earnedValuePerDay}</li>
					<li key="summaryEarnedValuePerWeek"><label>Earned Value Per Week:</label> {statistics.earnedValuePerWeek}</li>
				</ul>
				<h2>Tokens</h2>
				{statistics.tokens.map(function(token, index) {
					return (
						<div>
							<h3>{token.tokenName} ({token.tokenAddress})</h3>
							<ul>
								<li key="{token.tokenAddress}-firstTransactionAt"><label>First Transaction:</label> {token.firstTransactionAt}</li>
								<li key="{token.tokenAddress}-transationCount"><label>Transaction Count:</label> {token.transactionCount}</li>
								<li key="{token.tokenAddress}-tokenPrice"><label>Price:</label> {token.tokenPrice}</li>
								<li key="{token.tokenAddress}-tokenPriceSource"><label>Price Source:</label> {token.tokenPriceSource}</li>
								<li key="{token.tokenAddress}-tokenPriceUpdatedAt"><label>Price Updated At:</label> {token.tokenPriceUpdatedAt}</li>
								<li key="{token.tokenAddress}-tokenCount"><label>Count:</label> {token.tokenCount}</li>
								<li key="{token.tokenAddress}-value"><label>Value:</label> {token.value}</li>
								<li key="{token.tokenAddress}-earnedTokenCount"><label>Earned Count:</label> {token.earnedTokenCount}</li>
								<li key="{token.tokenAddress}-earnedTokenCountPerDay"><label>Earned Count Per Day:</label> {token.earnedTokenCountPerDay}</li>
								<li key="{token.tokenAddress}-earnedTokenCountPerWeek"><label>Earned Count Per Week:</label> {token.earnedTokenCountPerWeek}</li>
								<li key="{token.tokenAddress}-earnedValue"><label>Earned Value:</label> {token.earnedValue}</li>
								<li key="{token.tokenAddress}-earnedValuePerDay"><label>Earned Value Per Day:</label> {token.earnedValuePerDay}</li>
								<li key="{token.tokenAddress}-earnedValuePerWeek"><label>Earned Value Per Week:</label> {token.earnedValuePerWeek}</li>
							</ul>
						</div>
					)
				})}
			</div>;
		} else {
			console.log('render else');
			content = <span className="welcome">Welcome, please enter your wallet address</span>;
		}

		return (
			<div className="walletStatistics">
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

