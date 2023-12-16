
// ** UTILITY HELPOER FUNCTIONS **

// Function to format the release_date data from "YYYY-MM-DD" into "longMonth, numericDay, numericYear" format
export const formatReleaseDate = (dateString) => {
    const options = { year: 'numeric', month: 'short', day: 'numeric' };
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', options);
};

// Function to format the runtime data from "minutes" into "0h, 0m" format
export const formatRuntime = (minutes) => {
    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    return `${hours}h ${remainingMinutes}m`;
};

// Function to format the vote_count data from "integer" into "12.1k" if value was 12100
export const formatVoteCount = (voteCount) => {
    if (voteCount >= 1e9) {
        return `${voteCount >= 1e10 ? (voteCount / 1e9).toFixed(1) : (voteCount / 1e9).toFixed(2)}b votes`;
    } else if (voteCount >= 1e6) {
        return `${voteCount >= 1e7 ? (voteCount / 1e6).toFixed(1) : (voteCount / 1e6).toFixed(2)}m votes`;
    } else if (voteCount >= 1e3) {
        return `${voteCount >= 1e4 ? (voteCount / 1e3).toFixed(1) : (voteCount / 1e3).toFixed(2)}k votes`;
    } else {
        return `${voteCount} votes`;
    }
};

// Function to format the financial data of revenue & budget from "integer" into "$1.20m" if value was 1200000
export const formatFinancialData = (revenue) => {
    if (revenue >= 1e9) {
        return `$${revenue >= 1e10 ? (revenue / 1e9).toFixed(1) : (revenue / 1e9).toFixed(2)}b`;
    } else if (revenue >= 1e6) {
        return `$${revenue >= 1e7 ? (revenue / 1e6).toFixed(1) : (revenue / 1e6).toFixed(2)}m`;
    } else if (revenue >= 1e3) {
        return `$${revenue >= 1e4 ? (revenue / 1e3).toFixed(1) : (revenue / 1e3).toFixed(2)}k`;
    } else {
        return `$${revenue.toFixed(2)}`;
    }
};