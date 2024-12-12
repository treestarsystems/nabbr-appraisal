export interface Chart {
  appraisalId: string;
  memberInformation: MemberInformation;
  petInformation: PetInformation;
  appraisalInformation: AppraisalInformation;
}

export interface MemberInformation {
  name: string;
  memberNumber: string;
}

export interface PetInformation {
  name: string;
  age: number;
  dnaNumber: string;
  weight: number;
  color: string;
  markings: string;
  microchip: string;
  sex: string;
  registrationNumber: string;
}

export interface AppraisalInformation {
  mainDivision: MainDivision;
  appraiserName: string;
  appraiserNumber: string;
  seniorAppraiserName: string;
  seniorAppraiserNumber: string;
  date: string;
  additionalComments: string;
  place: number;
  appraisalScore: number;
}

export interface MainDivision {
  appearance: Division;
  head: Division;
  face: Division;
  neck: Division;
  forequarter: Division;
  centerPiece: Division;
  hindquarter: Division;
  skinCoat: Division;
  health: Division;
  temperament: Division;
  movement: Division;
}

export interface Division {
  characteristics: Characteristic[][];
  name: string;
  percentageWeight: number;
}

export interface Characteristic {
  name: string;
  score: string;
  value: number;
  factor: number;
  subDivision?: SubDivision;
}

export interface SubDivision {
  name: string;
}
