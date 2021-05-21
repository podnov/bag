import React, { Component } from 'react'
import {connect} from 'react-redux'
import equal from 'fast-deep-equal'
import Loader from 'react-loader-spinner'

import './WalletStatistics.css'

class WalletStatistics extends Component {
	constructor(props) {
		super(props);

		this.state = {
			hasError: this.props.hasError,
			isLoading: this.props.isLoading,
			statistics: this.props.statistics
		}
	}

	componentDidUpdate(prevProps) {
		if(!equal(this.props, prevProps)) {
			this.setState({
				hasError: this.props.hasError,
				isLoading: this.props.isLoading,
				statistics: this.props.statistics
			});
		}
	}


	render() {
		const hasError = this.state.hasError;
		const isLoading = this.state.isLoading;
		const statistics = this.state.statistics;

		let content;

		if (hasError) {
			content = <div>
				Error encountered fetching wallet data. Please verify wallet address and try again.
			</div>;
		} else if (isLoading) {
			content = <Loader type="MutatingDots" height={100} width={100} />;
		} else if (statistics) {
			let numberFormatter = new Intl.NumberFormat();
			let currencyFormatter = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' });

			let convertPriceSource = function(source) {
				let result;

				switch (source) {
					case 'PANCAKE_SWAP_V1':
						result = 'PancakeSwap v1';
						break;

					case 'PANCAKE_SWAP_V2':
						result = 'PancakeSwap v2';
						break;

					default:
						result = source;
				}

				return result;
			};

			statistics.tokens.sort(function(first, second) {
				return (first.tokenName + first.tokenAddress).toLowerCase().localeCompare(second.tokenName + second.tokenAddress);
			});

			content = <div>
				<h2>Summary</h2>
				<ul>
					<li key="summaryAddress"><label>Address:</label> {statistics.accountAddress}</li>
					<li key="summaryFirstTransactionAt"><label>First Transaction:</label> {new Date(statistics.firstTransactionAt).toLocaleString()}</li>
					<li key="summaryTransactionCount"><label>Transaction Count:</label> {numberFormatter.format(statistics.transactionCount)}</li>
					<li key="summaryValue"><label>Value:</label> {currencyFormatter.format(statistics.value)}</li>
					<li key="summaryAccruedValue"><label>Accrued Value:</label> {currencyFormatter.format(statistics.accruedValue)},&nbsp;
						<label>Per Day:</label> {currencyFormatter.format(statistics.accruedValuePerDay)},&nbsp;
						<label>Per Week:</label> {currencyFormatter.format(statistics.accruedValuePerWeek)}</li>
				</ul>
				<h2>Tokens</h2>
				{statistics.tokens.map(function(token, index) {
					return (
						<div>
							<h3>{token.tokenName} ({token.tokenAddress})</h3>
							<ul>
								<li key="{token.tokenAddress}-firstTransactionAt"><label>First Transaction:</label> {new Date(token.firstTransactionAt).toLocaleString()}</li>
								<li key="{token.tokenAddress}-transationCount"><label>Transaction Count:</label> {numberFormatter.format(token.transactionCount)}</li>
								<li key="{token.tokenAddress}-tokenPrice"><label>Price:</label> {token.tokenPrice}</li>
								<li key="{token.tokenAddress}-tokenPriceSource"><label>Price Source:</label> {convertPriceSource(token.tokenPriceSource)},&nbsp;
									<label>Updated At:</label> {new Date(token.tokenPriceUpdatedAt).toLocaleString()}</li>
								<li key="{token.tokenAddress}-tokenCount"><label>Count:</label> {numberFormatter.format(token.tokenCount)}</li>
								<li key="{token.tokenAddress}-value"><label>Value:</label> {currencyFormatter.format(token.value)}</li>
								<li key="{token.tokenAddress}-accruedTokenCount"><label>Accrued Count:</label> {numberFormatter.format(token.accruedTokenCount)},&nbsp;
									<label>Per Day:</label> {numberFormatter.format(token.accruedTokenCountPerDay)},&nbsp;
									<label>Per Week:</label> {numberFormatter.format(token.accruedTokenCountPerWeek)}</li>
								<li key="{token.tokenAddress}-accruedValue"><label>Accrued Value:</label> {currencyFormatter.format(token.accruedValue)},&nbsp;
									<label>Per Day:</label> {currencyFormatter.format(token.accruedValuePerDay)},&nbsp;
									<label>Per Week:</label> {currencyFormatter.format(token.accruedValuePerWeek)}</li>
							</ul>
						</div>
					)
				})}
			</div>;
		} else {
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
	hasError: state.wallet.hasError,
	isLoading: state.wallet.isLoading,
	statistics: state.wallet.statistics
})

export default connect(mapStateToProps)(WalletStatistics)

