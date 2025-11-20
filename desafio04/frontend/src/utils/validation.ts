import { z } from "zod";

// Schema de validação para CEP brasileiro
export const cepSchema = z
  .string()
  .min(1, "CEP é obrigatório")
  .regex(/^\d{5}-?\d{3}$/, "CEP deve estar no formato 00000-000")
  .transform((val) => val.replace(/\D/g, "")); // Remove hífen para enviar limpo

// Função para validar CEP
export function validateCep(cep: string): {
  success: boolean;
  error?: string;
  data?: string;
} {
  try {
    const cleanedCep = cepSchema.parse(cep);
    return { success: true, data: cleanedCep };
  } catch (error) {
    if (error instanceof z.ZodError) {
      return { success: false, error: error.issues[0].message };
    }
    return { success: false, error: "CEP inválido" };
  }
}

// Função para limpar CEP (apenas números)
export function cleanCep(cep: string): string {
  return cep.replace(/\D/g, "");
}
