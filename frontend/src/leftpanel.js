import React from 'react';

export class LeftPanel extends React.Component {
    renderLeaderBoard() {
        return (
            <LeaderBoard
                leaderBoard={this.props.leaderBoard}
                myrankline={this.props.myrankline}
                score={this.props.score}
                isLoggedIn={this.props.isLoggedIn}
            ></LeaderBoard>
        )
    }

    render() {
        return (
            <div className="leftpanel">
                <div>
                    LeaderBoard
                </div>
                {this.renderLeaderBoard()}
            </div>
        )
    }
}


class LeaderBoard extends React.Component {
    render() {
        let lastLine = '';
        if (this.props.isLoggedIn) {
            lastLine = (
                <tr>
                    <td> #{this.props.myrankline.rank} </td>
                    <td> {this.props.myrankline.nickName} </td>
                    <td> {this.props.score} </td>
                </tr>
            )
        }
        return (
            <div className='leaderboard'>
                <table>
                    <tbody>
                        {this.props.leaderBoard.map((rankLine, i) => {
                            return (
                                <tr key={i}>
                                    <td>
                                        #{rankLine.rank}
                                    </td>
                                    <td>
                                        {rankLine.nickName}
                                    </td>
                                    <td>
                                        {rankLine.score}
                                    </td>
                                </tr>
                            )
                        })}
                        {lastLine}
                    </tbody>
                </table>
            </div>
        )
    }
}