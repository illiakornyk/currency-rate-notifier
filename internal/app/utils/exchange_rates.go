package utils

// ConvertEURtoUSD,UAH takes the EUR to USD and EUR to UAH rates and converts them to a USD to UAH rate
func ConvertEURtoUSDUAH(eurToUSD, eurToUAH float64) float64 {
    return eurToUAH / eurToUSD
}
