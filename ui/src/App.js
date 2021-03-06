import React from 'react';
import logo from './logo.svg';

import Modal from 'react-modal';
import { Helmet } from 'react-helmet'

import './App.css';
import WalletInput from './components/wallet/WalletInput';
import WalletStatistics from './components/wallet/WalletStatistics';

function App() {

	Modal.setAppElement('*');
	const [ modalIsOpen, setModalIsOpen ] = React.useState(true);

	const customModalStyles = {
		content : {
			top : '50%',
			left : '50%',
			right : 'auto',
			bottom : 'auto',
			marginRight : '-50%',
			transform : 'translate(-50%, -50%)',
			maxHeight: '80vh'
		}
	};

	function closeModal(){
		setModalIsOpen(false);
	}

	function openModal(){
		setModalIsOpen(true);
	}

	return (
		<div className="App">
			<Helmet>
				<title>CryptoBag</title>
			</Helmet>
			<header>
				<h1>CryptoBag</h1>
				<div><button onClick={openModal}>Show Disclaimers</button></div>
				<WalletInput />
			</header>
			<div className="content">
				<WalletStatistics />
			</div>
			<footer>
				<div>
					Powered by:&nbsp;
					<a href="https://bscscan.com" rel="noreferrer" target="_blank">BscScan.com</a> APIs,&nbsp;
					<a href="https://github.com/pancakeswap/pancake-info-api" rel="noreferrer" target="_blank">Pancakeswap APIs</a>,&nbsp;
					<a href="https://0x.org/docs/api" rel="noreferrer" target="_blank">0x APIs</a>,&nbsp;
					<a href="https://github.com/podnov/bag" rel="noreferrer" target="_blank">GitHub</a>
				</div>
				© <a href="mailto:cryptobag.podnov@gmail.com" rel="noreferrer" target="_blank">CryptoBag</a> 2021. All Rights Reserved.
			</footer>

			<Modal
				isOpen={modalIsOpen}
				onRequestClose={closeModal}
				style={customModalStyles}
				contentLabel="Disclaimers"
				shouldCloseOnOverlayClick={false}
				>
				<h2>Disclaimers</h2>
				<ul>
					<li>The creator(s) of this website are not financial advisors and nothing on this site should be viewed as financial advice.</li>
					<li>Information provided on this website is provided on an "as is" basis without warranty of any kind, either express or implied. The materials and information on this website are intended for informational purposes only. Any information on this site may be out of date.</li>
					<li>Token past performance is not an indicator of future performance.</li>
					<li>This site is intended to be used with Binance Smart Chain wallets only.</li>
					<li>All monetary figures are shown in USD and are not intended to be used for any official purposes (i.e. taxes). All prices are provided by an external provider (PancakeSwap v1 or v2 at this time), and these are often delayed</li>
					<li>If you have over 10,000 transactions on your account, this site is NOT for you at this time due to API response size limits</li>
					<li>Accrued tokens and values refer to those tokens accrued through something like transaction tax-based reflection to holders. Accrued tokens and values calculations are likely to be skewed by recent transaction (&lt; 1hr old) due to transaction API lag. Accrued tokens are simplified as a calculation similar to (accrued = current - bought + sold) and may not be entirely accurate in all scenarios. Accrued tokens and values over time (per day or per week) are simplified calculations averaging accrued amount over the time since your first transaction for that token</li> 
					<li>By using this website you agree that the creator(s) cannot be held liable in respect to actions taken or not taken based on information contained on or missing from this website.</li>
				</ul>
				<div className="modalButtons"><button onClick={closeModal}>Agreed</button></div>
			</Modal>

		</div>
	);
}

export default App;
