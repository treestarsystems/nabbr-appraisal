export interface Chart {
  appraisalId: string;
  memberInformation: MemberInformation;
  petInformation: PetInformation;
  appraisalInformation: AppraisalInformation;
}

export interface MemberInformation {
  name: string;
  email: string;
  memberNumber: string;
}

export interface PetInformation {
  name: string;
  type: string;
  breed: string;
  age: string;
  dnaNumber: string;
  weight: string;
  color: string;
  markings: string;
  microchip: string;
  sex: string;
  registrationNumber: string;
}

export interface AppraisalInformation {
  mainDivision: MainDivision;
  appraiser: string;
  appraiserNumber: string;
  seniorAppraiser: string;
  seniorAppraiserNumber: string;
  date: string;
  value: string;
  notes: string;
  additionalComments: string;
  place: string;
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
  score: number;
  value: number;
  factor: number;
  subDivision?: SubDivision;
}

export interface SubDivision {
  name: string;
}
