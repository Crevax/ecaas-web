# ecaas
This project is for example purposes only. 
It is a mock service implementing a cost calculator for a moving service. 
It has been written to be compiled using `gomobile bind` for iOS and Android example SDKs

## Business Rules
This estimation service uses the following business rules when calculating an estimate

1. Multiply total estimated hours by the provided hourly rate for an initial subtotal
2. Apply the cost multiplier - effectively a service fee - to the subtotal
3. Apply tax to the subtotal plus service fee to generate the low end of the etimate
4. To estimate the high end of the estimate, if the preferred move date is on Friday or Saturday - add a complexity factor of 30% of the low estimate before tax - otherwise add a standard 15% weekday complexity factor to get a estimate high end subtotal
5. Apply tax to the high end of the estimate, end return the range formatted in USD

## Mobile App Examples

Android Version - [https://github.com/syllabix/ecaas-android](https://github.com/syllabix/ecaas-android)


iOS Version - [https://github.com/syllabix/ecaas-ios](https://github.com/syllabix/ecaas-ios)

